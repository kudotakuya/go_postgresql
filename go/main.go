package main

import (
	"go_api/controller"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/users", controller.GetUsers)
	e.POST("/users", controller.PostUsers)

	e.Logger.Fatal(e.Start(":8080"))
}
