package auth

import (
	"go-echo-sample/model"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var Config = echojwt.WithConfig(echojwt.Config{
	SigningKey: []byte("secret"),
})

func Signup(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	if user.Name == "" || user.Password == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid name or password",
		}
	}

	model.DB.Create(&user)
	user.Password = ""

	return c.JSON(http.StatusCreated, user)
}
