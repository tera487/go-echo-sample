package main

import (
	"go-echo-sample/controller"
	"go-echo-sample/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func connect(c echo.Context) error {
	db, _ := model.DB.DB()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		return c.String(http.StatusInternalServerError, "DB接続失敗しました")
	}

	return c.String(http.StatusOK, "DB接続しました")
}

func main() {
	e := echo.New()
	e.GET("/users", controller.GetUsers)
	e.GET("/users/:id", controller.GetUser)
	e.POST("/users", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)
	e.Logger.Fatal(e.Start(":8080"))
}
