package main

import (
	"fmt"
	"net/http"
	"path"

	"github.com/carlosghabrous/lenslocked/controllers"
	"github.com/carlosghabrous/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func main() {

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	t := views.Must(views.Parse(path.Join("templates", "home.gohtml")))
	router.Get("/", controllers.StaticHandler(t))

	t = views.Must(views.Parse(path.Join("templates", "contact.gohtml")))
	router.Get("/contact", controllers.StaticHandler(t))

	t = views.Must(views.Parse(path.Join("templates", "faq.gohtml")))
	router.Get("/faq", controllers.StaticHandler(t))

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", router)
}
