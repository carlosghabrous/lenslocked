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

	template := views.Must(views.Parse(path.Join("templates", "home.gohtml")))
	router.Get("/", controllers.StaticHandler(template))

	template = views.Must(views.Parse(path.Join("templates", "contact.gohtml")))
	router.Get("/contact", controllers.StaticHandler(template))

	template = views.Must(views.Parse(path.Join("templates", "faq.gohtml")))
	router.Get("/faq", controllers.StaticHandler(template))

	http.ListenAndServe(":3000", router)
}
