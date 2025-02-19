package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/carlosghabrous/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func executeTemplate(w http.ResponseWriter, templateName string) {
	templatePath := path.Join("templates", templateName+".gohtml")
	template, err := views.Parse(templatePath)
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal server error (template parsing)", http.StatusInternalServerError)
		return
	}

	template.Execute(w)

}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "home")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "contact")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "faq")
}

func getArticleById(w http.ResponseWriter, r *http.Request) {
	articleID := chi.URLParam(r, "articleID")
	// executeTemplate(w, "article")
	fmt.Fprintf(w, "<h1>Article ID: %s</h1>", articleID)
}

func main() {
	fmt.Println("starting server at port :3000")
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", homeHandler)
	router.Get("/contact", contactHandler)
	router.Get("/faq", faqHandler)

	router.Route("/articles", func(r chi.Router) {
		r.Route("/{articleID}", func(r chi.Router) {
			r.Get("/", getArticleById)
		})
	})

	http.ListenAndServe(":3000", router)
}
