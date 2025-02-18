package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func executeTemplate(w http.ResponseWriter, templateName string) {
	templatePath := path.Join("templates", templateName+".gohtml")
	template, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = template.Execute(w, nil)
	if err != nil {
		log.Printf("Error executingtemplate: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	executeTemplate(w, "home")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	executeTemplate(w, "contact")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	executeTemplate(w, "faq")
}

func getArticleById(w http.ResponseWriter, r *http.Request) {
	// articleID := chi.URLParam(r, "articleID")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	executeTemplate(w, "article")
	// fmt.Fprintf(w, "<h1>Article ID: %s</h1>", articleID)
}

func main() {
	fmt.Println("starting server at port :3000")
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", homeHandler)
	router.Get("/contact", contactHandler)
	router.Get("/faq", faqHandler)
	router.Get("/articles/{articleID}", getArticleById)
	http.ListenAndServe(":3000", router)
}
