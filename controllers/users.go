package controllers

import (
	"fmt"
	"net/http"

	"github.com/carlosghabrous/lenslocked/models"
)

type Users struct {
	Templates struct {
		New    Template
		SignIn Template
	}
	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.SignIn.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	user, err := u.UserService.Create(r.FormValue("email"), r.FormValue("password"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error creating user", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "user created with id: %+v", user)
}
