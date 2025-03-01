package controllers

import (
	"net/http"
)

func StaticHandler(template Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		template.Execute(w, nil)
	}
}

func Faq(template Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   string
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes, click on the banner",
		},
		{
			Question: "What are your support hours?",
			Answer:   "24/7"},
		{
			Question: "How do I cancel my subscription?",
			Answer:   `Email us at <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>`,
		},
		{
			Question: "Where is your office located?",
			Answer:   "We are located in the heart of downtown Toronto, Ontario, Canada"},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		template.Execute(w, questions)
	}
}
