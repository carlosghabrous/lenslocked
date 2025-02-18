package main

import (
	"html/template"
	"os"
)

type Person struct {
	Name      string
	Age       int
	Height    float32
	Phone     []int
	YearsCity map[int]string
	Bio       string
}

func main() {
	template, err := template.ParseFiles("hello.gohtml")

	if err != nil {
		panic(err)
	}

	carlos := Person{Name: "carlos", Age: 24, Height: 1.23, Phone: []int{1, 2, 3}, YearsCity: map[int]string{20: "Fuengirola", 15: "Geneva"}, Bio: "<script>alert('hacked')</script>"}

	err = template.ExecuteTemplate(os.Stdout, "hello.gohtml", carlos)
	if err != nil {
		panic(err)
	}
}
