package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if _, err := os.Open("non-existing-file"); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File does not exist error")
		} else {
			fmt.Println("Non-permission error")
		}
	}
}
