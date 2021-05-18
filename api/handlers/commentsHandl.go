package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GEtAllComments(c echo.Context) error {

	return c.JSON(http.StatusOK, "all comments... for now")
}
