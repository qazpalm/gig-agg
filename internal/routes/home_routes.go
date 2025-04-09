package routes 

import (
	"net/http"
	"github.com/qazpalm/gig-agg/internal/handlers/webhandlers"
	"github.com/qazpalm/gig-agg/internal/middleware"
)

// RegisterHomeRoutes registers the home routes with the given mux.
func RegisterHomeRoutes(mux *http.ServeMux, sessionMiddleware *middleware.SessionMiddleware) {
	homeHandler := webhandlers.NewHomeHandler(sessionMiddleware.GetSessionStore())
	
	mux.HandleFunc("/", homeHandler.HomeHandler)
	mux.HandleFunc("/about", homeHandler.AboutHandler)
}