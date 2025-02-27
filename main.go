package main

import (
	"log"
	"net/http"

	"github.com/carlosghabrous/lenslocked/controllers"
	"github.com/carlosghabrous/lenslocked/templates"
	"github.com/carlosghabrous/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	log.Println("starting server at port :3000")
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	template := views.Must(views.ParseFileSystem(templates.FileSystem, "home.gohtml", "layout-parts.gohtml"))
	router.Get("/", controllers.StaticHandler(template))

	template = views.Must(views.ParseFileSystem(templates.FileSystem, "contact.gohtml"))
	router.Get("/contact", controllers.StaticHandler(template))

	template = views.Must(views.ParseFileSystem(templates.FileSystem, "faq.gohtml"))
	router.Get("/faq", controllers.Faq(template))

	http.ListenAndServe(":3000", router)
}
