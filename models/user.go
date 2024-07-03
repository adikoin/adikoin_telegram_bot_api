package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                    primitive.ObjectID `json:"id" form:"id" bson:"_id,omitempty"`
	Balance               int                `json:"balance" bson:"balance" validate:"required"`
	LastFarmStartTime     time.Time          `json:"lastFarmStartTime" bson:"lastFarmStartTime"`
	Email                 string             `json:"email" bson:"email" validate:"required"`
	TelegramUserID        int64              `json:"telegramUserID" bson:"_telegramUserID"`
	IsBot                 bool               `json:"isBot" bson:"isBot"`
	FirstName             string             `json:"firstName" bson:"firstName" validate:"required"`
	LastName              string             `json:"lastName" bson:"lastName"`
	Username              string             `json:"username" bson:"username" validate:"required"`
	LanguageCode          string             `json:"languageCode" bson:"languageCode" validate:"required"`
	IsPremium             bool               `json:"isPremium" bson:"isPremium"`
	AddedToAttachmentMenu bool               `json:"addedToAttachmentMenu" bson:"addedToAttachmentMenu"`
	AllowsWriteToPm       bool               `json:"allowsWriteToPm" bson:"allowsWriteToPm"`
	PhotoUrl              string             `json:"photoUrl" bson:"photoUrl"`
	// CompanyID  primitive.ObjectID `json:"companyID" form:"companyID" bson:"_companyID,omitempty"`
	// *UserInput `bson:",inline"`
}

type UserEmailUpdate struct {
	Email string `json:"email" bson:"email" validate:"required"`
}

type UserUpdate struct {
	*UserInputUpdate `bson:",inline"`
	ID               primitive.ObjectID `json:"id" form:"id" bson:"_id,omitempty"`
}

type UserChangePassword struct {
	*ChangeUserPassword `bson:",inline"`
	ID                  primitive.ObjectID `json:"id" form:"id" bson:"_id,omitempty"`
}

type UserInput struct {
	FirstName string `json:"firstName" bson:"firstName" validate:"required"`
	LastName  string `json:"lastName" bson:"lastName" validate:"required"`
	Email     string `json:"email" bson:"email" validate:"required,email"`
	Password  string `json:"password,omitempty"  bson:"password" validate:"required"`
}

type UserInputUpdate struct {
	FirstName       string `json:"firstName" xml:"firstName" form:"firstName" bson:"firstName" validate:"required"`
	LastName        string `json:"lastName" xml:"lastName" form:"lastName" bson:"lastName" validate:"required"`
	CurrentPassword string `json:"currentPassword,omitempty" form:"currentPassword" xml:"currentPassword,omitempty" bson:"currentPassword" validate:"required"`
	CompanyCategory string `json:"companyCategory" bson:"companyCategory"`
}

type ChangeUserPassword struct {
	CurrentPassword string `json:"currentPassword,omitempty" form:"currentPassword" xml:"currentPassword,omitempty" bson:"currentPassword" validate:"required"`
	Password        string `json:"password" form:"password" xml:"password" bson:"password" validate:"required"`
}

type LoginInput struct {
	Email string `json:"email" bson:"email" validate:"required,email"`
	// Phone    string `json:"phone" form:"phone" bson:"phone" validate:"required"`
	Password string `json:"password" form:"password" bson:"password" validate:"required"`
}

type UserDropDown struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName string             `json:"firstName" bson:"firstName"`
}

type AllDropDownUsers struct {
	Data []UserDropDown `json:"data" xml:"data"`
}

type Task struct {
	ID                   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title                string             `json:"title" bson:"title"`
	Subtitle             string             `json:"subtitle" bson:"subtitle"`
	Icon                 string             `json:"icon" bson:"icon"`
	RewardAmount         int                `json:"rewardAmount" bson:"rewardAmount"`
	Host                 string             `json:"host" bson:"host"`
	Path                 string             `json:"path" bson:"path"`
	StartedTelegramIDs   []int64            `json:"startedTelegramIDs" bson:"startedTelegramIDs"`
	CompletedTelegramIDs []int64            `json:"completedTelegramIDs" bson:"completedTelegramIDs"`
	ClaimedTelegramIDs   []int64            `json:"claimedTelegramIDs" bson:"claimedTelegramIDs"`
	Status               string             `json:"status" bson:"status"`
	Type                 string             `json:"type" bson:"type"`
}

type Tasks struct {
	Data []Task `json:"data"`
}

// type Subscribe struct {
// ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
// Email string             `json:"email" form:"email" bson:"email" validate:"required"`
// *SubscribeInput `bson:",inline"`
//
// Status     bool               `json:"status" bson:"status"`
// UpdatedAt  time.Time          `json:"updatedAt" bson:"updatedAt" validate:"required"`
// CreatedAt  time.Time          `json:"createdAt" bson:"createdAt" validate:"required"`
// }
