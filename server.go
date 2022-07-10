package main

import (
	db "myapp1/db"
	u "myapp1/users"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db.DB()
	db.Migrate()

	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, `{"response" : "success"}`)
	})

	e.GET("/users", u.GetUsers)
	e.GET("/users/:id", u.GetUser)
	e.POST("/users", u.Save)
	e.PUT("/users/:id", u.Update)
	e.DELETE("/users/:id", u.Delete)

	e.Logger.Fatal(e.Start(":1234"))
}
