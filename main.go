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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tpl, err := template.ParseFiles(path.Join("templates", templateName+".gohtml"))
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "error parsing the "+templateName+"template", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "error executing the"+templateName+"template", http.StatusInternalServerError)
		return
	}

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "home")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "contact")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>FAQ page</h1><ul><li>Q: first question</li><li>A: first answer</li></ul>")
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", homeHandler)
	router.Get("/contact", contactHandler)
	router.Get("/faq", faqHandler)
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", router)
}
