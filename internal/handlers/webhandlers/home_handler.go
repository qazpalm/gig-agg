package webhandlers

import (
	"net/http"
	"fmt"

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
	
	// Get the session token from the cookie
	cookie, err := r.Cookie("session_token")
	if err != nil {
		// No session token found, try to get the remember token
		rememberToken, err := r.Cookie("remember_token")
		if err == nil {
			// TODO: Check if rememberToken is valid
			// If valid, create a new session token and set it in the cookie
			_ = rememberToken
		} else {
			// No session token or remember token found, user is not logged in
			// Redirect to login page or show a message
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
	}
	sessionToken := cookie.Value
	session, exists := hh.sessionStore.GetSession(sessionToken)
	fmt.Println("Session Token:", sessionToken)
	fmt.Println("Session:", session)
	if !exists {
		// Session not found, user is not logged in
		// Redirect to login page or show a message
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	

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