package http

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/haffjjj/myblog-backend/models"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

//AuthHandler ...
type AuthHandler struct{}

//NewAuthHandler ...
func NewAuthHandler(e *echo.Echo) {
	handler := &AuthHandler{}
	e.POST("/auth", handler.Auth)
}

//Auth ...
func (a *AuthHandler) Auth(c echo.Context) error {
	claims := &models.JWTClaims{
		"hafizjoundy",
		"hafizjoundy@codex.works",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := viper.GetString("jwtSecret")

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
