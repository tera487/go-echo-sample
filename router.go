package main

import (
	"errors"
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
	e.POST("/signup", auth.Signup) // POST /signup

	api := e.Group("/api")
	api.Use(auth.Config)
	api.GET("/", func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
		if !ok {
			return errors.New("JWT token missing or invalid")
		}
		claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
		if !ok {
			return errors.New("failed to cast claims as jwt.MapClaims")
		}
		return c.JSON(http.StatusOK, claims)
	})

	// api.DELETE("/users/:id", controller.DeleteUser)
	return e
}
