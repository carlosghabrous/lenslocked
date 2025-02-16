package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Hello World 2c</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>email me <a href=\"mailto:carlos.ghabrous@gmail.com\">here</a></p>")
}

type Router struct{}

func (*Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.NotFound(w, r)

	}
}
func main() {
	// http.HandleFunc("/", pathHandler)

	fmt.Println("starting server at port :3000")
	myHandler := &Router{}

	http.ListenAndServe(":3000", myHandler)
}
