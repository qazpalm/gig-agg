package routes 

import (
	"net/http"
	"github.com/qazpalm/gig-agg/internal/handlers/webhandlers"
	"github.com/qazpalm/gig-agg/internal/middleware"
	"github.com/qazpalm/gig-agg/internal/store"
	"github.com/qazpalm/gig-agg/internal/session"
	"github.com/qazpalm/gig-agg/internal/auth"
)

// RegisterHomeRoutes registers the home routes with the given mux.
func RegisterHomeRoutes(mux *http.ServeMux, sessionMiddleware *middleware.SessionMiddleware, userStore store.UserStore, sessionStore *session.SessionStore, userAuthManager *auth.UserAuthManager) {
	homeHandler := webhandlers.NewHomeHandler(sessionMiddleware.GetSessionStore())
	userHandler := webhandlers.NewUserHandler(userStore, userAuthManager, sessionStore)
	
	mux.HandleFunc("/", homeHandler.HomeHandler)
	mux.HandleFunc("/about", homeHandler.AboutHandler)

	mux.HandleFunc("GET /login", userHandler.LoginHandler)
	mux.HandleFunc("POST /login", userHandler.LoginHandler)

	mux.HandleFunc("GET /create-account", userHandler.CreateAccountHandler)
	mux.HandleFunc("POST /create-account", userHandler.CreateAccountHandler)
}