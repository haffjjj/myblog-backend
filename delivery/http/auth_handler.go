package http

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/haffjjj/myblog-backend/models"
	"github.com/haffjjj/myblog-backend/usecase/auth"
	"github.com/spf13/viper"

	"github.com/labstack/echo"
)

//AuthHandler ...
type AuthHandler struct {
	aUsecase auth.Usecase
}

//NewAuthHandler ...
func NewAuthHandler(c *echo.Echo, aU auth.Usecase) {
	handler := &AuthHandler{aU}
	c.POST("/auth", handler.Auth)
}

//Credential ...
type Credential struct {
	Username string
	Password string
}

//Auth is method Authhandler
func (a *AuthHandler) Auth(c echo.Context) error {

	var ct Credential
	c.Bind(&ct)

	fmt.Println(ct)

	user, err := a.aUsecase.Auth(ct.Username, ct.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	claims := &models.JWTClaims{
		user.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := viper.GetString("jwtSecret")

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
