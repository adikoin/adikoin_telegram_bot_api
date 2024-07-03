package routes

import (
	"telegram_bot_api/controller"

	"github.com/labstack/echo/v4"
)

func GetUserApiRoutes(e *echo.Echo, userController *controller.UserController) {
	v1 := e.Group("/api/v1")
	{
		v1.GET("/users/:id/check_email", userController.CheckIsRegistered)
		v1.GET("/users/:id", userController.GetUser)
		v1.POST("/users/:id/start_farm", userController.StartFarm)
		v1.POST("/users/:id/claim", userController.Claim)
		v1.PUT("/users/:id/edit_email", userController.UpdateUserEmail)
		v1.GET("/users/:id/tasks", userController.GetUserTasks)
		v1.POST("/users/:id/tasks/:task_id/start_task", userController.StartTask)
		v1.POST("/users/:id/tasks/:task_id/check_task", userController.CheckTask)
		v1.POST("/users/:id/tasks/:task_id/claim_task", userController.ClaimTask)
		// v1.POST("/login", userController.AuthenticateUser)
		// v1.GET("/dropdownusers", userController.GetDropdownUsers)
		// v1.POST("/signup", userController.SaveUser)

		// v1.DELETE("/users/:id", userController.DeleteUser)
		// e.GET("/dashboard", userController.HomePage)
		// v1.PUT("/users/:id/changepassword", userController.ChangeUserPassword)
	}
}
