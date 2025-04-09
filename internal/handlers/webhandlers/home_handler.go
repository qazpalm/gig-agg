package webhandlers

import (
	"net/http"

	"github.com/qazpalm/gig-agg/internal/session"
)

type HomeHandler struct {
	sessionStore *session.SessionStore
}

func NewHomeHandler(sessionStore *session.SessionStore) *HomeHandler {
	return &HomeHandler{
		sessionStore: sessionStore,
	}
}

// HomeHandler handles the home page request.
func (hh *HomeHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	
	sessionToken := r.Header.Get("session_token")
	session, exists := hh.sessionStore.GetSession(sessionToken)


	data := struct {
		Title string
		IsLoggedIn bool
		Username string
	} {
		Title: "Home - Gig-Agg",
		IsLoggedIn: exists,
		Username: session.Username,
	}
	
	RenderTemplate(w, "home.html", data)
}

func (hh *HomeHandler) AboutHandler(w http.ResponseWriter, r *http.Request) {
	sessionToken := r.Header.Get("session_token")
	session, exists := hh.sessionStore.GetSession(sessionToken)

	data := struct {
		Title string
		IsLoggedIn bool
		Username string
	} {
		Title: "Home - Gig-Agg",
		IsLoggedIn: exists,
		Username: session.Username,
	}
	RenderTemplate(w, "about.html", data)
}