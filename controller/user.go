package controller

import (
	"net/http"
	"strconv"
	model "telegram_bot_api/models"
	"telegram_bot_api/repository"
	"telegram_bot_api/security"
	"telegram_bot_api/util"

	"github.com/labstack/echo/v4"
)

// const (
// 	farmDuration   = 8 * time.Hour
// 	pointsPerCycle = 80
// )

type UserController struct {
	userRepository repository.UserRepository
	authValidator  *security.AuthValidator
}

func NewUserController(userRepository repository.UserRepository, authValidator *security.AuthValidator) *UserController {
	return &UserController{userRepository: userRepository, authValidator: authValidator}
}

func (userController *UserController) UpdateUserEmail(c echo.Context) error {

	id := c.Param("id")

	telegramUserID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	payload := new(model.UserEmailUpdate)

	if err := util.BindAndValidate(c, payload); err != nil {
		return err
	}

	err = userController.userRepository.UpdateUserEmail(telegramUserID, payload)
	if err != nil {
		return util.Negotiate(c, http.StatusNotFound, nil)
	}

	return util.Negotiate(c, http.StatusCreated, nil)
}

func (userController *UserController) GetUser(c echo.Context) error {

	id := c.Param("id")

	telegramUserID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	user, err := userController.userRepository.FindUserByTelegramID(telegramUserID)

	if err != nil {
		return err
	}

	// log.Println(user.LastFarmStartTime)

	return util.Negotiate(c, http.StatusOK, user)

}

func (userController *UserController) CheckIsRegistered(c echo.Context) error {

	id := c.Param("id")

	telegramUserID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	user, err := userController.userRepository.FindUserByTelegramID(telegramUserID)

	if err != nil {
		return err
	}

	if user.Email == "" {
		return util.Negotiate(c, http.StatusNotFound, user.Email)
	} else {
		return util.Negotiate(c, http.StatusOK, user.Email)
	}
}

func (userController *UserController) StartFarm(c echo.Context) error {

	id := c.Param("id")

	if id == "" {
		util.Negotiate(c, http.StatusNotFound, nil)
	}

	telegramUserID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	err = userController.userRepository.StartFarm(telegramUserID)

	if err != nil {
		return err
	}

	return util.Negotiate(c, http.StatusCreated, nil)
}

func (userController *UserController) Claim(c echo.Context) error {

	id := c.Param("id")

	if id == "" {
		util.Negotiate(c, http.StatusNotFound, nil)
	}

	telegramUserID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	err = userController.userRepository.Claim(telegramUserID)

	if err != nil {
		return nil
	}

	return util.Negotiate(c, http.StatusCreated, nil)
}

func (userController *UserController) GetUserTasks(c echo.Context) error {

	id := c.Param("id")

	telegramUserID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	tasks, err := userController.userRepository.GetUserTasks(telegramUserID)

	if err != nil {
		return err
	}

	return util.Negotiate(c, http.StatusOK, tasks)
}

func (userController *UserController) StartTask(c echo.Context) error {

	id := c.Param("id")
	taskID := c.Param("task_id")

	if id == "" {
		util.Negotiate(c, http.StatusNotFound, nil)
	}

	telegramUserID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	err = userController.userRepository.StartTask(telegramUserID, taskID)

	if err != nil {
		return err
	}

	return util.Negotiate(c, http.StatusCreated, nil)
}

func (userController *UserController) CheckTask(c echo.Context) error {

	id := c.Param("id")
	taskID := c.Param("task_id")

	if id == "" {
		util.Negotiate(c, http.StatusNotFound, nil)
	}

	telegramUserID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	err = userController.userRepository.CheckTask(telegramUserID, taskID)

	if err != nil {
		return err
	}

	return util.Negotiate(c, http.StatusCreated, nil)
}

func (userController *UserController) ClaimTask(c echo.Context) error {

	id := c.Param("id")
	taskID := c.Param("task_id")

	if id == "" {
		util.Negotiate(c, http.StatusNotFound, nil)
	}

	telegramUserID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	err = userController.userRepository.ClaimTask(telegramUserID, taskID)

	if err != nil {
		return err
	}

	return util.Negotiate(c, http.StatusCreated, nil)
}

// func (userController *UserController) SaveUser(c echo.Context) error {

// 	payload := new(model.UserInput)

// 	if err := util.BindAndValidate(c, payload); err != nil {
// 		return err
// 	}

// 	_, err := userController.userRepository.FindByEmail(payload.Email)
// 	if err == nil {
// 		return exception.ConflictException("User", "email", payload.Email)
// 	}

// func (userController *UserController) NewSubscribe(c echo.Context) error {

// 	payload := new(model.Subscribe)

// 	if err := util.BindAndValidate(c, payload); err != nil {
// 		return err
// 	}

// 	_, err := userController.userRepository.FindSubscribeByEmail(payload.Email)
// 	if err == nil {
// 		// return nil
// 		return exception.ConflictException("Email", "email", payload.Email)
// 	}

// 	// subscribe := &model.Subscribe{SubscribeInput: payload}
// 	_, err = userController.userRepository.SaveSubscribe(payload)
// 	if err != nil {
// 		return err
// 	}

// 	return util.Negotiate(c, http.StatusCreated, "Email saved")
// }

// func (userController *UserController) FarmStatus(c echo.Context) error {

// 	// id := c.Param("id")

// 	// telegramUserID, err := strconv.ParseInt(id, 10, 64)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// _, err = userController.userRepository.FindSubscribeByID(telegramUserID)
// 	// if err == nil {
// 	// 	// return nil
// 	// 	balance, err := userController.userRepository.FindUserBalanceByTelegramID(telegramUserID)
// 	// 	if err != nil {
// 	// 		return util.Negotiate(c, http.StatusNotFound, "StatusNotFound")
// 	// 	}
// 	// 	return util.Negotiate(c, http.StatusFound, balance)
// 	// }

// 	return util.Negotiate(c, http.StatusNotFound, "StatusNotFound")
// }

// func (userController *UserController) AuthenticateUser(c echo.Context) error {

// 	payload := new(model.LoginInput)
// 	if err := util.BindAndValidate(c, payload); err != nil {
// 		return err
// 	}

// 	// log.Println("end time is:", payload.Password)

// 	user, valid := userController.authValidator.ValidateCredentials(payload.Email, payload.Password)
// 	if !valid {
// 		return exception.UnauthorizedException()
// 	}

// 	jwt, err := util.GenerateJwtToken(user)
// 	if err != nil {
// 		return err
// 	}

// 	user.Password = ""

// 	return util.Negotiate(c, http.StatusOK, model.Token{Token: jwt})
// }

// func (userController *UserController) SaveUser(c echo.Context) error {

// 	payload := new(model.UserInput)

// 	if err := util.BindAndValidate(c, payload); err != nil {
// 		return err
// 	}

// 	_, err := userController.userRepository.FindByEmail(payload.Email)
// 	if err == nil {
// 		return exception.ConflictException("User", "email", payload.Email)
// 	}

// 	user := &model.User{UserInput: payload}

// 	//encrypt password
// 	err = beforeSave(user)
// 	if err != nil {
// 		return err
// 	}

// 	createdUser, err := userController.userRepository.SaveUser(user)
// 	if err != nil {
// 		return err
// 	}

// 	return util.Negotiate(c, http.StatusCreated, createdUser)
// }

// func (userController *UserController) GetAllUser(c echo.Context) error {
// 	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)
// 	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 64)

// 	pagedUser, _ := userController.userRepository.GetAllUser(page, limit)
// 	return util.Negotiate(c, http.StatusOK, pagedUser)
// }

// func (userController *UserController) UpdateUser(c echo.Context) error {

// 	id := c.Param("id")

// 	payload := new(model.UserInputUpdate)

// 	if err := util.BindAndValidate(c, payload); err != nil {
// 		return err
// 	}

// 	currentUser, err := userController.userRepository.FindById(id)

// 	if err != nil || util.VerifyPassword(currentUser.Password, payload.CurrentPassword) != nil {
// 		return util.Negotiate(c, http.StatusUnauthorized, nil)
// 	}

// 	updateduser, err := userController.userRepository.UpdateUser(id, &model.UserUpdate{UserInputUpdate: payload})
// 	if err != nil {
// 		return util.Negotiate(c, http.StatusNotFound, updateduser)
// 	}

// 	return util.Negotiate(c, http.StatusOK, updateduser)
// }

// func (userController *UserController) DeleteUser(c echo.Context) error {
// 	id := c.Param("id")

// 	err := userController.userRepository.DeleteUser(id)
// 	if err != nil {
// 		return err
// 	}
// 	return c.NoContent(http.StatusNoContent)
// }

// func beforeSave(user *model.User) (err error) {
// 	hashedPassword, err := util.EncryptPassword(user.Password)
// 	if err != nil {
// 		return err
// 	}
// 	user.Password = string(hashedPassword)
// 	return nil
// }

// func (userController *UserController) HomePage(c echo.Context) error {
// 	// jwt, _ := c.Cookie("jwt")
// 	// log.Println("jwt is:", jwt)
// 	// c.Response().Header().Set("Authorization", "Bearer "+jwt.Value)
// 	// c.Response().WriteHeader(201)
// 	return c.Render(http.StatusOK, "dashboard", nil)

// }

// func (userController *UserController) ChangeUserPassword(c echo.Context) error {

// 	id := c.Param("id")

// 	payload := new(model.ChangeUserPassword)

// 	if err := util.BindAndValidate(c, payload); err != nil {
// 		return err
// 	}

// 	currentUser, err := userController.userRepository.FindById(id)

// 	if err != nil || util.VerifyPassword(currentUser.Password, payload.CurrentPassword) != nil {
// 		return util.Negotiate(c, http.StatusUnauthorized, nil)
// 	}

// 	hashedPassword, err := util.EncryptPassword(payload.Password)
// 	if err != nil {
// 		return err
// 	}
// 	payload.Password = string(hashedPassword)

// 	changeduserpassword, err := userController.userRepository.ChangeUserPassword(id, &model.UserChangePassword{ChangeUserPassword: payload})
// 	if err != nil {
// 		return util.Negotiate(c, http.StatusNotFound, changeduserpassword)
// 	}

// 	return util.Negotiate(c, http.StatusOK, changeduserpassword)
// }

// func (userController *UserController) GetDropdownUsers(c echo.Context) error {
// 	users, _ := userController.userRepository.GetDropdownUsers()
// 	return util.Negotiate(c, http.StatusOK, users)
// }
