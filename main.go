package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello World 2c</h1>")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("starting server at port :3000")
	http.ListenAndServe(":3000", nil)
}
