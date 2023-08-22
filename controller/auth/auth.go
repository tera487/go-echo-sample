package auth

import (
	"fmt"
	"go-echo-sample/model"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var Config = echojwt.WithConfig(echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return new(jwtCustomClaims)
	},
	SigningKey: []byte("secret"),
})

// 独自クレーム型
type jwtCustomClaims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

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

type LoginFormValue struct {
	Name     string `json:"name" `
	Password string `json:"password" `
}

func Login(c echo.Context) error {
	lfv := new(LoginFormValue)
	if err := c.Bind(lfv); err != nil {
		return err
	}

	u := new(model.User)
	if err := model.DB.Where("name = ?", lfv.Name).First(&u).Error; err != nil {
		return fmt.Errorf("get users by email failed , %w", err)
	}

	if u.ID == 0 || u.Password != lfv.Password {
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "invalid name or password",
		}
	}

	// 独自クレーム作成
	claims := &jwtCustomClaims{
		strconv.FormatUint(uint64(u.ID), 10),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	// トークン作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return fmt.Errorf("login failed at NewWithClaims err %w", err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
