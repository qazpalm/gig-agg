package webhandlers

import (
	"fmt"
	"net/http"
	"html/template"
)

func RenderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	tmpl, err := template.ParseFiles(
		"./assets/templates/base.html",
		"./assets/templates/title_bar.html",
		"./assets/templates/" + templateName,
	)
	if err != nil {
		fmt.Printf("Error loading template %s: %v", templateName, err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Printf("Error executing template %s: %v", templateName, err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}