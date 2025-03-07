package main

import (
	"fmt"
	"os"
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
	fmt.Printf("TODO: hash password %q\n", password)
}

func compare(password, hash string) {
	fmt.Printf("TODO: compare password %q with hash %q\n", password, hash)
}
