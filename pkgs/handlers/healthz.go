package handlers

import (
	"github.com/labstack/echo/v5"
)

func Healthz(c echo.Context) error {
	return c.JSON(200, "OK")
}
