package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	//pw_crypt := "$2b$04$DGsFgFEoetT7F3LEHOsLmOBam5GTmEBPByTFWCq1DP3dD15iJFNeC"

	password := []byte("123")
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println("Encrypted Pass", string(hashedPassword))

	match := CheckPasswordHash("123", string(hashedPassword))
	fmt.Println(match)
}
