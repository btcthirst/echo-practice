package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {

	return c.JSON(http.StatusOK, "all users... for now")
}
