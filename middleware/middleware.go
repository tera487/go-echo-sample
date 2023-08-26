package middleware

import (
	"errors"
	"go-echo-sample/controller/auth"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return errors.New("JWT token missing or invalid")
		}

		_, err := auth.JwtParser(token)
		if err != nil {
			return err
		}

		if err := next(c); err != nil {
			return err
		}

		return nil
	}
}
