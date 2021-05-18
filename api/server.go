package api

import (
	"github.com/btcthirst/echo-practice/api/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunServ(port string) {

	//start echo
	e := echo.New()

	api := e.Group("/api/v1")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api.GET("/", handlers.Home)
	//admin
	adm := api.Group("/admin")

	adm.GET("/", handlers.AdmiGo, middleware.BasicAuth(basic))

	//users
	u := api.Group("/users")

	u.GET("/", handlers.GetAllUsers)

	//posts
	p := api.Group("/posts")

	p.GET("/", handlers.GetAllPosts)

	//comments
	c := api.Group("/comments")

	c.GET("/", handlers.GEtAllComments)

	//start server
	e.Logger.Fatal(e.Start(port))
}

func basic(uname, pass string, c echo.Context) (bool, error) {

	if uname == "maxdev" && pass == "mustChange" {
		return true, nil
	}
	return false, nil
}
