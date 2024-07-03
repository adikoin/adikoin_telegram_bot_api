package security

import (
	"telegram_bot_api/config"
	model "telegram_bot_api/models"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var whiteListPaths = []string{
	"/favicon.ico",
	"/static*",
	"/",
	"/api/v1/users/:id/check_email",
	"/api/v1/users/:id/start_farm",
	"/api/v1/users/:id/edit_email",
	"/api/v1/users/:id",
	"/api/v1/users/:id/claim",
	"/api/v1/users/:id/tasks",
	"/api/v1/users/:id/tasks/:task_id/start_task",
	"/api/v1/users/:id/tasks/:task_id/check_task",
	"/api/v1/users/:id/tasks/:task_id/claim_task",
}

// change default error message
// func init() {
// 	middleware.ErrJWTMissing.Code = 401
// 	middleware.ErrJWTMissing.Message = "Unauthorized"
// }

func WebSecurityConfig(e *echo.Echo) {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.JwtCustomClaims)
		},
		SigningKey: []byte(config.JWTSecret),
		Skipper:    skipAuth,
	}

	// config := middleware.JWTConfig{
	// 	Claims:     &model.JwtCustomClaims{},
	// 	SigningKey: []byte(config.JWTSecret),
	// 	Skipper:    skipAuth,
	// }

	// config := echojwt.Config{
	// 	NewClaimsFunc:     &model.JwtCustomClaims{},
	// 	SigningKey: []byte(config.JWTSecret),
	// 	Skipper:    skipAuth,
	// 	SigningKey: []byte("secret"),
	// }

	e.Use(echojwt.WithConfig(config))
	// e.Use(middleware.JWTWithConfig(config))
}

func skipAuth(c echo.Context) bool {
	// log.Println(c.Path())
	// Skip authentication for and signup login requests
	for _, path := range whiteListPaths {
		if path == c.Path() {

			return true
		}
	}
	return false
}
