package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RedirectIndexPage(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/")
}
