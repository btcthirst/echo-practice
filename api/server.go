package api

import (
	"github.com/btcthirst/echo-practice/api/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunServ(p string) {

	//start echo
	e := echo.New()

	api := e.Group("/api/v1")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api.GET("/", handlers.Home)

	//start server
	e.Logger.Fatal(e.Start(p))
}
