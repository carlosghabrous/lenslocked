package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}

	return t
}

func Parse(templatePath string) (Template, error) {
	tpl, err := template.ParseFiles(templatePath)

	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}

	return Template{htmlTpl: tpl}, nil
}

type Template struct {
	htmlTpl *template.Template
}

func (t *Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := t.htmlTpl.Execute(w, nil)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "error executing the template", http.StatusInternalServerError)
		return
	}

}