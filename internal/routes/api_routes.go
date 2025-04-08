package routes

import (
	"net/http"
	"github.com/qazpalm/gig-agg/internal/handlers/apihandlers"
	"github.com/qazpalm/gig-agg/internal/store"
)

func RegisterAPIRoutes(mux *http.ServeMux, artistStore store.ArtistStore, genreStore store.GenreStore) {
	artistHandler 	:= apihandlers.NewArtistHandler(artistStore)
	genreHandler 	:= apihandlers.NewGenreHandler(genreStore)

	mux.HandleFunc("POST /api/artist", 		artistHandler.CreateArtist)
	mux.HandleFunc("GET /api/artist/{id}", 	artistHandler.GetArtist)
	mux.HandleFunc("GET /api/artist", 		artistHandler.GetArtists)
	mux.HandleFunc("PUT /api/artist/{id}", 	artistHandler.UpdateArtist)
	mux.HandleFunc("DELETE /api/artist/{id}", artistHandler.DeleteArtist)

	mux.HandleFunc("POST /api/genre", 		genreHandler.CreateGenre)
	mux.HandleFunc("GET /api/genre/{id}", 	genreHandler.GetGenre)
	mux.HandleFunc("GET /api/genre", 		genreHandler.GetGenres)
	mux.HandleFunc("PUT /api/genre/{id}", 	genreHandler.UpdateGenre)
	mux.HandleFunc("DELETE /api/genre/{id}", genreHandler.DeleteGenre)
}

