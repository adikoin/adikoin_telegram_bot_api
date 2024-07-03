package config

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CORSConfig(e *echo.Echo) {
	config := middleware.CORSConfig{
		AllowOrigins: []string{"*", "https://www.chatmaster.kz"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderAccessControlAllowOrigin, echo.HeaderCacheControl, echo.HeaderConnection},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}
	e.Use(middleware.CORSWithConfig(config))
}
