package webhandlers

import (
	"net/http"
)

// HomeHandler handles the home page request.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.html", nil)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about.html", nil)
}