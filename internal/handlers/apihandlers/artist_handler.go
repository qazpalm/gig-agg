package apihandlers

import (
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/qazpalm/gig-agg/internal/store"
	"github.com/qazpalm/gig-agg/internal/models"
)

// ArtistHandler handles requests related to artists.
type artistHandler struct {
	store store.ArtistStore
}

// NewArtistHandler creates a new artist handler.
func NewArtistHandler(store store.ArtistStore) *artistHandler {
	return &artistHandler{store: store}
}

// CreateArtist handles the creation of a new artist.
func (h *artistHandler) CreateArtist(w http.ResponseWriter, r *http.Request) {
	newArtist := &models.Artist{}
	
	// Get the artist data from the request body
	err := json.NewDecoder(r.Body).Decode(newArtist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the artist data
	if newArtist.Name == "" || newArtist.Description == "" || newArtist.SpotifyID == "" {
		http.Error(w, "Invalid artist data", http.StatusBadRequest)
		return
	}

	// Create the artist in the database
	_, err = h.store.CreateArtist(newArtist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetArtist handles the retrieval of an artist by ID.
func (h *artistHandler) GetArtist(w http.ResponseWriter, r *http.Request) {
	idPath := r.PathValue("id")
	id, err := strconv.Atoi(idPath)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	artist, err := h.store.GetArtist(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if artist == nil {
		http.Error(w, "Artist not found", http.StatusNotFound)
		return
	}

	// Return the artist data as JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(artist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


