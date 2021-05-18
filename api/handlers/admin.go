package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AdmiGo(c echo.Context) error {

	return c.JSON(http.StatusOK, "all you want U can do... for now")
}
