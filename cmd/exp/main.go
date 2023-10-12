package main

import (
	"html/template"
	"os"
)

type User struct {
	Name    string
	Age     int
	Weight  float64
	Lottery []int
	Phones  map[string]int
	Meta    UserMeta
}

type UserMeta struct {
	Visits      int
	Conditional bool
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	// user := User{Name: "Carlos", Bio: `<script>alert("Ups");</script>`, Age: 123}
	user := User{
		Name: "carlos", Age: 453,
		Weight: 68.2, Lottery: []int{1, 56, 26, 93, 2},
		Phones: map[string]int{"carlos": 41456633},
		Meta:   UserMeta{Visits: 4, Conditional: false}}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
