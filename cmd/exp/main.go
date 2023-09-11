package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string
	Age  int
	Meta UserMeta
}

type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{Name: "Carlos", Bio: `<script>alert("Ups");</script>`, Age: 123}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
