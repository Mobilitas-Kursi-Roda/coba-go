package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type RegisterSchema struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func main() {
	validare := validator.New()
	err := validare.Struct(&RegisterSchema{
		Email:    "test@mail.com",
		Username: "test",
	})
	if err != nil {
		fmt.Println()
	}
}
