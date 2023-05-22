package controller

import (
	"go-echo-sample/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateUser は
func CreateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	model.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}
