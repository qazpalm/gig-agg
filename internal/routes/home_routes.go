package routes 

import (
	"net/http"
	"github.com/qazpalm/gig-agg/internal/handlers"
)

// RegisterHomeRoutes registers the home routes with the given mux.
func RegisterHomeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.HomeHandler)
}