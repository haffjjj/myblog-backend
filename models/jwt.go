package models

import jwt "github.com/dgrijalva/jwt-go"

//JWTClaims ...
type JWTClaims struct {
	Username string
	jwt.StandardClaims
}
