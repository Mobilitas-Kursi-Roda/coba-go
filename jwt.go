package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

type UserClaims struct {
	Id    string `json:"id"`
	First string `json:"first"`
	Last  string `json:"last"`
	jwt.StandardClaims
}

func NewAccessToken(claims UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

func main() {

	a := UserClaims{
		Id:    "1",
		First: "John",
		Last:  "Doe",
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}

	signedAccessToken, _ := NewAccessToken(a)

	fmt.Println(signedAccessToken)

}
