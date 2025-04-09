package webhandlers

import (
	"github.com/qazpalm/gig-agg/internal/auth"
	"github.com/qazpalm/gig-agg/internal/session"
	"github.com/qazpalm/gig-agg/internal/store"
	"net/http"
)

// UserHandler is a struct that handles user-related HTTP requests.
type UserHandler struct {
	userStore     	store.UserStore
	userAuthManager *auth.UserAuthManager
	sessionStore  	*session.SessionStore
}

// NewUserHandler creates a new UserHandler with the given user store and session store.
func NewUserHandler(userStore store.UserStore, authManager *auth.UserAuthManager, sessionStore *session.SessionStore) *UserHandler {
	return &UserHandler{
		userStore:      userStore,
		userAuthManager: authManager,
		sessionStore:   sessionStore,
	}
}

func (userHandler *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		_, err := userHandler.userAuthManager.LoginUser(email, password, w)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Redirect to the home page after successful login
		http.Redirect(w, r, "/", http.StatusSeeOther)

		return
	}

	if r.Method == http.MethodGet {
		data := struct {
			Title string
			IsLoggedIn bool
			Username string
		} {
			Title: "Login - Gig-Agg",
			IsLoggedIn: false,
			Username: "",
		}
		RenderTemplate(w, "login.html", data)
	}
}

func (userHandler *UserHandler) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		username := r.FormValue("username")
		password := r.FormValue("password")

		_, err := userHandler.userAuthManager.RegisterUser(email, username, password)
		if err != nil {
			http.Error(w, "Error creating account", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodGet {
		data := struct {
			Title string
			IsLoggedIn bool
			Username string
		} {
			Title: "Create Account - Gig-Agg",
			IsLoggedIn: false,
			Username: "",
		}
		RenderTemplate(w, "create_account.html", data)
	}
}