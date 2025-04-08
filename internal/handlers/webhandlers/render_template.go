package webhandlers

import (
	"net/http"
	"html/template"
)

func RenderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	tmpl, err := template.ParseFiles(
		"./assets/templates/" + templateName,
		"./assets/templates/title_bar.html",
	)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}