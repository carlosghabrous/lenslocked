package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func usage() {
	fmt.Println("Usage:\tbcrypt hash <password>")
	fmt.Println("\tbcrypt compare <password> <hash>")
	os.Exit(1)
}

func main() {
	switch os.Args[1] {
	case "hash":
		if len(os.Args) < 3 {
			usage()
		}
		hash(os.Args[2])
	case "compare":
		if len(os.Args) < 4 {
			usage()
		}
		compare(os.Args[2], os.Args[3])
	default:
		usage()
	}
}

func hash(password string) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("error while hashing password: %v\n", err)
	}
	fmt.Println(string(hashedBytes))
}

func compare(hash, password string) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Println("password does not match")
		return
	}
	fmt.Println("password matches")
}
