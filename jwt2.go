package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

var secret = []byte("e7462a4f5295b5001cdb93eb3d6c65775910324ce38faacdf9e19403f4a3ca43")

func main() {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["created_at"] = time.Now()
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = "maxim"

	tokenString, _ := token.SignedString(secret)

	fmt.Println(tokenString)

}
