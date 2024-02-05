package mian

import "errors"

func mian() {
	err := Createuser()
}

func Connect() error {
	return errors.New("connection failed")
}

func Createuser() error {
	err := Connect()
	if err != nil {
		return err
	}
	// ...
	return nil
}
