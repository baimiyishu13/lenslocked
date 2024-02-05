package main

import (
	"errors"
	"fmt"
)

func main() {
	err := CreateOrg()
	fmt.Print(err)
}

func Connect() error {
	return errors.New("connection failed")
}

func Createuser() error {
	err := Connect()
	if err != nil {
		// return err
		return fmt.Errorf("create user: %w", err)
	}
	// ...
	return nil
}

func CreateOrg() error {
	err := Createuser()
	if err != nil {
		return fmt.Errorf("creaete org: %w", err)
	}
	return nil
}
