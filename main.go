package main

import (
	"log"
	"net/http"
	"path"

	"github.com/carlosghabrous/lenslocked/controllers"
	"github.com/carlosghabrous/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	log.Println("starting server at port :3000")
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// parse templates
	template, err := views.Parse(path.Join("templates", "home.gohtml"))
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	router.Get("/", controllers.StaticHandler(template))

	template, err = views.Parse(path.Join("templates", "contact.gohtml"))
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	router.Get("/contact", controllers.StaticHandler(template))

	template, err = views.Parse(path.Join("templates", "faq.gohtml"))
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	router.Get("/faq", controllers.StaticHandler(template))

	http.ListenAndServe(":3000", router)
}
