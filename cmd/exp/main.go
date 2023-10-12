package main

import (
	"errors"
	"fmt"
)

func main() {
	err := CreateUser()
	fmt.Println(err)
}

func Connect() error {
	return errors.New("could not connect")
}

func CreateUser() error {
	err := Connect()
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}
	// ...

	return nil
}