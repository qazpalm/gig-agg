package routes

import (
	"net/http"
	"github.com/qazpalm/gig-agg/internal/handlers/apihandlers"
	"github.com/qazpalm/gig-agg/internal/store"
	"github.com/qazpalm/gig-agg/internal/middleware"
	"github.com/qazpalm/gig-agg/internal/apikeys"
)

func RegisterAPIRoutes(mux *http.ServeMux, artistStore store.ArtistStore, genreStore store.GenreStore, venueStore store.VenueStore, apiKeyManager *apikeys.APIKeyManager) {
	// Middleware for API key authentication
	apiKeyMiddleware := middleware.NewAPIKeyMiddleware(apiKeyManager)

	artistHandler 	:= apihandlers.NewArtistHandler(artistStore)
	genreHandler 	:= apihandlers.NewGenreHandler(genreStore)
	venueHandler 	:= apihandlers.NewVenueHandler(venueStore)

	// Register API routes with middleware
	mux.HandleFunc("POST /api/artist", apiKeyMiddleware.ServeAuthorised(http.HandlerFunc(artistHandler.CreateArtist)).ServeHTTP)
	mux.HandleFunc("GET /api/artist/{id}", apiKeyMiddleware.ServeAuthorised(http.HandlerFunc(artistHandler.GetArtist)).ServeHTTP)
	mux.HandleFunc("GET /api/artist", apiKeyMiddleware.ServeAuthorised(http.HandlerFunc(artistHandler.GetArtists)).ServeHTTP)
	mux.HandleFunc("PUT /api/artist/{id}", apiKeyMiddleware.ServeAuthorised(http.HandlerFunc(artistHandler.UpdateArtist)).ServeHTTP)
	mux.HandleFunc("DELETE /api/artist/{id}", apiKeyMiddleware.ServeAuthorised(http.HandlerFunc(artistHandler.DeleteArtist)).ServeHTTP)

	mux.HandleFunc("POST /api/genre", apiKeyMiddleware.ServeAuthorised(http.HandlerFunc(genreHandler.CreateGenre)).ServeHTTP)
	mux.HandleFunc("GET /api/genre/{id}", apiKeyMiddleware.ServeAuthorised(http.HandlerFunc(genreHandler.GetGenre)).ServeHTTP)
	mux.HandleFunc("GET /api/genre", apiKeyMiddleware.ServeAuthorised(http.HandlerFunc(genreHandler.GetGenres)).ServeHTTP)
	mux.HandleFunc("PUT /api/genre/{id}", apiKeyMiddleware.ServeAuthorised(http.HandlerFunc(genreHandler.UpdateGenre)).ServeHTTP)
	mux.HandleFunc("DELETE /api/genre/{id}", apiKeyMiddleware.ServeAuthorised(http.HandlerFunc(genreHandler.DeleteGenre)).ServeHTTP)

	mux.HandleFunc("POST /api/venue", apiKeyMiddleware.ServeAuthorised(http.HandlerFunc(venueHandler.CreateVenue)).ServeHTTP)
	mux.HandleFunc("GET /api/venue/{id}", apiKeyMiddleware.ServeAuthorised(http.HandlerFunc(venueHandler.GetVenue)).ServeHTTP)
	mux.HandleFunc("GET /api/venue", apiKeyMiddleware.ServeAuthorised(http.HandlerFunc(venueHandler.GetVenues)).ServeHTTP)
	mux.HandleFunc("PUT /api/venue/{id}", apiKeyMiddleware.ServeAuthorised(http.HandlerFunc(venueHandler.UpdateVenue)).ServeHTTP)
	mux.HandleFunc("DELETE /api/venue/{id}", apiKeyMiddleware.ServeAuthorised(http.HandlerFunc(venueHandler.DeleteVenue)).ServeHTTP)


}

