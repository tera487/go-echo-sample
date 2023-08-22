package main

import (
	"errors"
	"fmt"
	"go-echo-sample/controller"
	"go-echo-sample/controller/auth"
	"go-echo-sample/model"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func newRouter() *echo.Echo {
	e := echo.New()

	model.SetupDB()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", controller.GetUsers)
	e.GET("/users/:id", controller.GetUser)
	e.POST("/users", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser)

	e.POST("/signup", auth.Signup)
	e.POST("/login", auth.Login)

	api := e.Group("/api")
	api.Use(auth.Config)
	api.GET("/", func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return errors.New("JWT token missing or invalid")
		}

		token, _ = jwt.Parse(token.Raw, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return auth.Config, nil
		})

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok && token.Valid {
			return errors.New("failed to cast claims as jwt.MapClaims")
		}
		return c.JSON(http.StatusOK, claims)
	})

	// api.DELETE("/users/:id", controller.DeleteUser)
	return e
}
