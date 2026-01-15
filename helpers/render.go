package helpers

import (
	"html/template"
	"net/http"
)

func Render(w http.ResponseWriter, page string, data map[string]interface{}) {
	files := []string{
		"views/layouts/base.html",
		"views/layouts/navbar.html",
		"views/" + page,
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
