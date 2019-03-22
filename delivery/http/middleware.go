package http

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

//Middleware represent middleware struct
type Middleware struct{}

//JWTAuth is method from middleware
func (m *Middleware) JWTAuth() echo.MiddlewareFunc {
	secret := viper.GetString("jwtSecret")

	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(secret),
	})
}
