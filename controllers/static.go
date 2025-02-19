package controllers

import (
	"net/http"

	"github.com/carlosghabrous/lenslocked/views"
)

func StaticHandler(template views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		template.Execute(w)
	}
}
