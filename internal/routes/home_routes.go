package routes 

import (
	"net/http"
	"github.com/qazpalm/gig-agg/internal/handlers"
	"github.com/qazpalm/gig-agg/internal/middleware"
)

// RegisterHomeRoutes registers the home routes with the given mux.
func RegisterHomeRoutes(mux *http.ServeMux, sessionMiddleware *middleware.SessionMiddleware) {
	mux.HandleFunc("/", sessionMiddleware.ServeSessionProtected(http.HandlerFunc(handlers.HomeHandler)).ServeHTTP)
}