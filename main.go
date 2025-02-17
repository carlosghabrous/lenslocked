package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Hello World 2c</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact</h1><p>Contact me <a href=\"mailto:carlos.ghabrous@gmail.com\">here</a></p>")
}

func getArticleById(w http.ResponseWriter, r *http.Request) {
	articleID := chi.URLParam(r, "articleID")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Article ID: %s</h1>", articleID)
}

func main() {
	fmt.Println("starting server at port :3000")
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", homeHandler)
	router.Get("/contact", contactHandler)
	router.Get("/articles/{articleID}", getArticleById)
	http.ListenAndServe(":3000", router)
}
