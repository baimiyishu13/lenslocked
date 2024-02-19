package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// for i, arg := range os.Args {
	// 	fmt.Println(i, ":", arg)
	// }
	switch os.Args[1] {
	case "hash":
		hash(os.Args[2])
	case "compaer":
		compaer(os.Args[2], os.Args[3])
	default:
		fmt.Println("Invalid command")
	}
}

func hash(password string) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Failed to hash the password: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(hashedBytes))
}

func compaer(password, hash string) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Printf("Failed to compare the password: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Password is correct")
}
