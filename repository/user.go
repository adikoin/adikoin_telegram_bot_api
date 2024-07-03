package repository

import (
	"context"
	"errors"
	"telegram_bot_api/exception"
	model "telegram_bot_api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var cntx context.Context = context.TODO()

// const (
// 	farmDuration   = 6 * time.Hour
// 	pointsPerCycle = 80
// )

type UserRepository interface {
	UpdateUserEmail(id int64, user *model.UserEmailUpdate) error
	GetUserTasks(id int64) (*model.Tasks, error)
	// SaveSubscribe(subscribe *model.Subscribe) (*model.Subscribe, error)
	// FindSubscribeByEmail(email string) (*model.Subscribe, error)
	// FindSubscribeByID(id int64) (*model.Subscribe, error)
	FindUserByTelegramID(id int64) (*model.User, error)
	StartFarm(id int64) error
	Claim(id int64) error
	StartTask(id int64, taskID string) error
	CheckTask(id int64, taskID string) error
	ClaimTask(id int64, taskID string) error
	// GetAllUser(page int64, limit int64) (*model.PagedUser, error)
	// GetDropdownUsers() (*model.AllDropDownUsers, error)
	// SaveUser(user *model.User) (*model.User, error)
	// FindByEmail(email string) (*model.User, error)
	// FindByPhone(phone string) (*model.User, error)
	// FindById(id string) (*model.User, error)
	// UpdateUser(id string, user *model.UserUpdate) (*model.UserUpdate, error)
	// DeleteUser(id string) error
	// ChangeUserPassword(id string, user *model.UserChangePassword) (*model.UserChangePassword, error)
}

type userRepositoryImpl struct {
	Connection *mongo.Database
}

func NewUserRepository(Connection *mongo.Database) UserRepository {
	return &userRepositoryImpl{Connection: Connection}
}

func (userRepository *userRepositoryImpl) UpdateUserEmail(id int64, email *model.UserEmailUpdate) error {

	filter := bson.M{"_telegramUserID": id}
	update := bson.M{
		"$set": bson.M{
			"email": email.Email,
		},
	}

	result, err := userRepository.Connection.Collection("users").UpdateOne(cntx, filter, update)

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return err
	}

	return nil
}

func (userRepository *userRepositoryImpl) FindUserByTelegramID(id int64) (*model.User, error) {
	var existingUser model.User
	filter := bson.M{"_telegramUserID": id}
	err := userRepository.Connection.Collection("users").FindOne(cntx, filter).Decode(&existingUser)
	if err != nil {
		return nil, err
	}
	return &existingUser, nil
}

func (userRepository *userRepositoryImpl) StartFarm(id int64) error {

	var existingUser model.User
	filter := bson.M{"_telegramUserID": id}
	err := userRepository.Connection.Collection("users").FindOne(cntx, filter).Decode(&existingUser)
	if err != nil {
		return err
	}

	now := time.Now().Local()

	existingUser.LastFarmStartTime = now

	// update := bson.M{
	// 	"$set": bson.M{
	// 		"firstName": user.FirstName,
	// 		"lastName":  user.LastName,
	// 	},
	// }

	result, err := userRepository.Connection.Collection("users").UpdateOne(cntx, bson.M{"_telegramUserID": id}, bson.M{"$set": existingUser}, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return exception.ResourceNotFoundException("User", "id", "id")
	}

	return nil
}

func (userRepository *userRepositoryImpl) Claim(id int64) error {

	var existingUser model.User
	filter := bson.M{"_telegramUserID": id}
	err := userRepository.Connection.Collection("users").FindOne(cntx, filter).Decode(&existingUser)
	if err != nil {
		return err
	}

	now := time.Now().UTC()

	duration := now.Sub(existingUser.LastFarmStartTime)

	if duration.Hours() >= 8 {
		existingUser.Balance = existingUser.Balance + 80

		result, err := userRepository.Connection.Collection("users").UpdateOne(cntx, bson.M{"_telegramUserID": id}, bson.M{"$set": existingUser}, options.Update().SetUpsert(true))
		if err != nil {
			return err
		}
		if result.MatchedCount == 0 {
			return exception.ResourceNotFoundException("User", "id", "id")
		}
	}

	return nil
}

func (userRepository *userRepositoryImpl) GetUserTasks(id int64) (*model.Tasks, error) {

	filter := bson.M{}
	// projection := bson.D{
	// 	{Key: "startedTelegramIDs", Value: 0},
	// 	{Key: "completedTelegramIDs", Value: 0},
	// 	{Key: "claimedTelegramIDs", Value: 0},
	// }

	collection := userRepository.Connection.Collection("tasks")

	cursor, err := collection.Find(cntx, filter)
	if err != nil {
		return nil, err
	}

	var tasks []model.Task

	if err = cursor.All(cntx, &tasks); err != nil {
		return nil, err
	}

	for i, task := range tasks {

		if contains(task.StartedTelegramIDs, id) {
			tasks[i].Status = "started"
		}

		if contains(task.CompletedTelegramIDs, id) {
			tasks[i].Status = "completed"
		}

		if contains(task.ClaimedTelegramIDs, id) {
			tasks[i].Status = "claimed"
		}

		tasks[i].StartedTelegramIDs = nil
		tasks[i].CompletedTelegramIDs = nil
		tasks[i].ClaimedTelegramIDs = nil

	}

	return &model.Tasks{
		Data: tasks,
	}, nil

}

func contains(slice []int64, item int64) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func (userRepository *userRepositoryImpl) StartTask(id int64, taskID string) error {

	objectId, _ := primitive.ObjectIDFromHex(taskID)
	filter := bson.M{"_id": objectId, "startedTelegramIDs": bson.M{"$ne": id}}
	update := bson.M{"$addToSet": bson.M{"startedTelegramIDs": id}}

	result, err := userRepository.Connection.Collection("tasks").UpdateOne(cntx, filter, update)

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return exception.ResourceNotFoundException("User", "id", "id")
	}

	return nil
}

func (userRepository *userRepositoryImpl) CheckTask(id int64, taskID string) error {

	// var existingTask model.Task

	// objectId, _ := primitive.ObjectIDFromHex(taskID)

	// filter := bson.M{"_id": objectId}

	// err := userRepository.Connection.Collection("users").FindOne(cntx, filter).Decode(&existingTask)

	// if err != nil {
	// 	return err
	// }

	// filter := bson.M{"_id": objectId, "startedTelegramIDs": bson.M{"$ne": id}}
	// update := bson.M{"$addToSet": bson.M{"startedTelegramIDs": id}}

	// result, err := userRepository.Connection.Collection("tasks").UpdateOne(cntx, filter, update)

	// if err != nil {
	// 	return err
	// }

	// if result.MatchedCount == 0 {
	// 	return exception.ResourceNotFoundException("User", "id", "id")
	// }

	objectId, _ := primitive.ObjectIDFromHex(taskID)
	filter := bson.M{"_id": objectId, "completedTelegramIDs": bson.M{"$ne": id}}
	update := bson.M{"$addToSet": bson.M{"completedTelegramIDs": id}}

	result, err := userRepository.Connection.Collection("tasks").UpdateOne(cntx, filter, update)

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return exception.ResourceNotFoundException("User", "id", "id")
	}

	return nil
}

func (userRepository *userRepositoryImpl) ClaimTask(id int64, taskID string) error {

	// var existingUser model.User
	// filter := bson.M{"_telegramUserID": id}
	// err := userRepository.Connection.Collection("users").FindOne(cntx, filter).Decode(&existingUser)
	// if err != nil {
	// 	return err
	// }

	objectId, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id":                  objectId,
		"completedTelegramIDs": bson.M{"$in": []int64{id}},
		"claimedTelegramIDs":   bson.M{"$ne": id},
	}

	update := bson.M{"$addToSet": bson.M{"claimedTelegramIDs": id}}

	var task model.Task
	err = userRepository.Connection.Collection("tasks").FindOneAndUpdate(cntx, filter, update).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("condition not met or task not found")
		}
		return err
	}

	userFilter := bson.M{"_telegramUserID": id}
	userUpdate := bson.M{"$inc": bson.M{"balance": task.RewardAmount}}

	result, err := userRepository.Connection.Collection("users").UpdateOne(cntx, userFilter, userUpdate)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}

// func (userRepository *userRepositoryImpl) SaveSubscribe(subscribe *model.Subscribe) (*model.Subscribe, error) {

// 	_, err := userRepository.Connection.Collection("subscribe").InsertOne(cntx, subscribe)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return subscribe, nil
// }

// func (userRepository *userRepositoryImpl) FindSubscribeByEmail(email string) (*model.Subscribe, error) {
// 	var existingSubscribe model.Subscribe
// 	filter := bson.M{"email": email}
// 	err := userRepository.Connection.Collection("subscribe").FindOne(cntx, filter).Decode(&existingSubscribe)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &existingSubscribe, nil
// }

// func (userRepository *userRepositoryImpl) FindSubscribeByID(id int64) (*model.Subscribe, error) {
// 	var existingSubscribe model.Subscribe
// 	filter := bson.M{"_telegramUserID": id}
// 	err := userRepository.Connection.Collection("subscribe").FindOne(cntx, filter).Decode(&existingSubscribe)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &existingSubscribe, nil
// }

// func (userRepository *userRepositoryImpl) FindByEmail(email string) (*model.User, error) {
// 	var existingUser model.User
// 	filter := bson.M{"email": email}
// 	err := userRepository.Connection.Collection("users").FindOne(cntx, filter).Decode(&existingUser)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &existingUser, nil
// }

// func (userRepository *userRepositoryImpl) FindByPhone(phone string) (*model.User, error) {
// 	var existingUser model.User
// 	filter := bson.M{"phone": phone}
// 	err := userRepository.Connection.Collection("users").FindOne(cntx, filter).Decode(&existingUser)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &existingUser, nil
// }

// func (userRepository *userRepositoryImpl) FindById(id string) (*model.User, error) {
// 	var existingUser model.User
// 	objectId, _ := primitive.ObjectIDFromHex(id)
// 	filter := bson.M{"_id": objectId}
// 	err := userRepository.Connection.Collection("users").FindOne(cntx, filter).Decode(&existingUser)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &existingUser, nil
// }

// func (userRepository *userRepositoryImpl) SaveUser(user *model.User) (*model.User, error) {
// 	user.ID = primitive.NewObjectID()

// 	_, err := userRepository.Connection.Collection("users").InsertOne(cntx, user)
// 	if err != nil {
// 		return nil, err
// 	}

// 	user.Password = ""
// 	return user, nil
// }

// func (userRepository *userRepositoryImpl) UpdateUser(id string, user *model.UserUpdate) (*model.UserUpdate, error) {

// 	objectId, _ := primitive.ObjectIDFromHex(id)

// 	filter := bson.M{"_id": objectId}
// 	update := bson.M{
// 		"$set": bson.M{
// 			"firstName": user.FirstName,
// 			"lastName":  user.LastName,
// 		},
// 	}

// 	result, err := userRepository.Connection.Collection("users").UpdateMany(cntx, filter, update)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if result.MatchedCount == 0 {
// 		return nil, exception.ResourceNotFoundException("User", "id", id)
// 	}

// 	user.ID = objectId
// 	user.CurrentPassword = ""
// 	return user, nil
// }

// func (userRepository *userRepositoryImpl) DeleteUser(id string) error {
// 	objectId, _ := primitive.ObjectIDFromHex(id)
// 	filter := bson.M{"_id": objectId}

// 	result, err := userRepository.Connection.Collection("users").DeleteOne(cntx, filter)
// 	if err != nil {
// 		return err
// 	}
// 	if result.DeletedCount == 0 {
// 		return exception.ResourceNotFoundException("User", "id", id)
// 	}

// 	return nil
// }

// func (userRepository *userRepositoryImpl) ChangeUserPassword(id string, user *model.UserChangePassword) (*model.UserChangePassword, error) {

// 	objectId, _ := primitive.ObjectIDFromHex(id)

// 	filter := bson.M{"_id": objectId}
// 	update := bson.M{
// 		"$set": bson.M{
// 			"password": user.Password,
// 		},
// 	}

// 	result, err := userRepository.Connection.Collection("users").UpdateOne(cntx, filter, update)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if result.MatchedCount == 0 {
// 		return nil, exception.ResourceNotFoundException("User", "id", id)
// 	}

// 	user.ID = objectId
// 	user.CurrentPassword = ""
// 	return user, nil
// }

// func (userRepository *userRepositoryImpl) GetAllUser(page int64, limit int64) (*model.PagedUser, error) {
// 	var users []model.User

// 	filter := bson.M{}

// 	collection := userRepository.Connection.Collection("users")

// 	//	projection := bson.D{
// 	//		{"id", 1},
// 	//		{"firstName", 1},
// 	//		{"lastName", 1},
// 	//		{"email", 1},
// 	//	}
// 	//

// 	projection := bson.D{
// 		{
// 			Key:   "id",
// 			Value: 1,
// 		},
// 		{
// 			Key:   "firstName",
// 			Value: 1,
// 		}, {
// 			Key:   "lastName",
// 			Value: 1,
// 		},
// 		{
// 			Key:   "email",
// 			Value: 1,
// 		}}

// 	paginatedData, err := paginate.New(collection).Context(cntx).Limit(limit).Page(page).Select(projection).Filter(filter).Decode(&users).Find()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &model.PagedUser{
// 		Data:     users,
// 		PageInfo: paginatedData.Pagination,
// 	}, nil
// }

// func (userRepository *userRepositoryImpl) GetDropdownUsers() (*model.AllDropDownUsers, error) {

// 	filter := bson.M{}

// 	collection := userRepository.Connection.Collection("users")

// 	cursor, err := collection.Find(cntx, filter)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var users []model.UserDropDown

// 	if err = cursor.All(cntx, &users); err != nil {
// 		log.Fatal(err)
// 	}

// 	return &model.AllDropDownUsers{
// 		Data: users,
// 	}, nil

// }
