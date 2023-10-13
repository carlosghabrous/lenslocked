package controllers

import (
	"net/http"

	"github.com/carlosghabrous/lenslocked/views"
)

func StaticHandler(t views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nil)
	}
}