package controller

import (
	"go_api/model"

	"net/http"

	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	query := new(model.User)
	if err := c.Bind(query); err != nil {
		return err
	}
	result := model.GetUsers(query)
	return c.JSON(http.StatusOK, result)
}

func PostUsers(c echo.Context) error {
	query := new(model.User)
	if err := c.Bind(query); err != nil {
		return err
	}
	model.CreateUser(query)
	return c.JSON(http.StatusOK, "success")
}
