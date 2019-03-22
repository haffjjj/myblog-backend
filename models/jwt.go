package models

import jwt "github.com/dgrijalva/jwt-go"

//JWTClaims represent model for jwtClaims data
type JWTClaims struct {
	Username string
	jwt.StandardClaims
}
