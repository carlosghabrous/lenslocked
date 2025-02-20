package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

type Template struct {
	htmlTemplate *template.Template
}

func Must(template Template, err error) Template {
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	return template
}

func ParseFileSystem(fileSystem fs.FS, patterns ...string) (Template, error) {
	parsedTemplate, err := template.ParseFS(fileSystem, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}

	return Template{
		htmlTemplate: parsedTemplate,
	}, nil
}

func (t Template) Execute(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := t.htmlTemplate.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
