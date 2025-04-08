package routes

import (
	"net/http"
	"github.com/qazpalm/gig-agg/internal/handlers/apihandlers"
	"github.com/qazpalm/gig-agg/internal/store"
	"github.com/qazpalm/gig-agg/internal/middleware"
	"github.com/qazpalm/gig-agg/internal/apikeys"
)

func RegisterAPIRoutes(mux *http.ServeMux, artistStore store.ArtistStore, genreStore store.GenreStore, apiKeyManager *apikeys.APIKeyManager) {
	// Middleware for API key authentication
	apiKeyMiddleware := middleware.NewAPIKeyMiddleware(apiKeyManager)

	artistHandler 	:= apihandlers.NewArtistHandler(artistStore)
	genreHandler 	:= apihandlers.NewGenreHandler(genreStore)

	// Register API routes with middleware
	mux.Handle("POST /api/artist", apiKeyMiddleware.ServeHTTP(artistHandler.CreateArtist))
	mux.Handle("GET /api/artist/{id}", apiKeyMiddleware.ServeHTTP(artistHandler.GetArtist))
	mux.Handle("GET /api/artist", apiKeyMiddleware.ServeHTTP(artistHandler.GetArtists))
	mux.Handle("PUT /api/artist/{id}", apiKeyMiddleware.ServeHTTP(artistHandler.UpdateArtist))
	mux.Handle("DELETE /api/artist/{id}", apiKeyMiddleware.ServeHTTP(artistHandler.DeleteArtist))

	mux.HandleFunc("POST /api/genre", 		genreHandler.CreateGenre)
	mux.HandleFunc("GET /api/genre/{id}", 	genreHandler.GetGenre)
	mux.HandleFunc("GET /api/genre", 		genreHandler.GetGenres)
	mux.HandleFunc("PUT /api/genre/{id}", 	genreHandler.UpdateGenre)
	mux.HandleFunc("DELETE /api/genre/{id}", genreHandler.DeleteGenre)
}

