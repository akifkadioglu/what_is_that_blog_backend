package models

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	Name string `json:"name"`
	Id   uint   `json:"id"`
	Time string `json:"time"`
	jwt.StandardClaims
}
