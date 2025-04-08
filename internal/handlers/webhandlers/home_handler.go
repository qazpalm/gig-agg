package webhandlers

import (
	"net/http"
)

// HomeHandler handles the home page request.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
		IsLoggedIn bool
	} {
		Title: "Home - Gig-Agg",
		IsLoggedIn: false,
	}
	RenderTemplate(w, "home.html", data)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
		IsLoggedIn bool
	} {
		Title: "About - Gig-Agg",
		IsLoggedIn: false,
	}
	RenderTemplate(w, "about.html", data)
}