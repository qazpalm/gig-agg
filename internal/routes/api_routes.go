package routes

import (
	"net/http"
	"github.com/qazpalm/gig-agg/internal/handlers/apihandlers"
	"github.com/qazpalm/gig-agg/internal/store"
)

func RegisterAPIRoutes(mux *http.ServeMux, store store.ArtistStore) {
	handler := apihandlers.NewArtistHandler(store)

	mux.HandleFunc("POST /api/artist", handler.CreateArtist)
	mux.HandleFunc("GET /api/artist/{id}", handler.GetArtist)
}