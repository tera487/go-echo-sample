package main

import (
	"go-echo-sample/controller"
	"go-echo-sample/controller/auth"
	authMiddleware "go-echo-sample/middleware"
	"go-echo-sample/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func newRouter() *echo.Echo {
	e := echo.New()

	model.SetupDB()

	e.Use(middleware.Logger(), middleware.Recover())

	// Custom Error Handler
	e.HTTPErrorHandler = customHTTPErrorHandler

	e.GET("/users", controller.GetUsers)
	e.GET("/users/:id", controller.GetUser)
	e.POST("/users", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser)

	e.POST("/signup", auth.Signup)
	e.POST("/login", auth.Login)

	api := e.Group("/api")
	api.Use(auth.Config)
	api.Use(authMiddleware.AuthMiddleware)
	api.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "success!!")
	})

	// api.DELETE("/users/:id", controller.DeleteUser)
	return e
}

func customHTTPErrorHandler(err error, c echo.Context) {
	c.Logger().Error(err)

	if err := c.JSON(http.StatusInternalServerError, map[string]string{
		"error": err.Error(),
	}); err != nil {
		c.Logger().Error(err)
	}
}
