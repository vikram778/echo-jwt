package model

import jwt "github.com/dgrijalva/jwt-go"

//Token struct declaration
type Token struct {
	UserID   uint
	Email    string
	UserName string
	*jwt.StandardClaims
}
