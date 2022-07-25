package model

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string
	Email    string
	jwt.StandardClaims
}

var JwtSecretKey = []byte("secretKey")

func VerifyToken(token string) bool {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtSecretKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("ErrSignatureInvalid")
			return false
		}
		fmt.Println("time is over")
		return false
	}
	if !tkn.Valid {
		fmt.Println("Invalid: ")
		return false
	}
	return true
}
