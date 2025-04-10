package apihandlers

import (
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/qazpalm/gig-agg/internal/store"
	"github.com/qazpalm/gig-agg/internal/models"
)

type getGenresBody struct {
	Count  int `json:"count"`
	Offset int `json:"offset"`
}

// GenreHandler handles requests related to genres.
type GenreHandler struct {
	store store.GenreStore
}

// NewGenreHandler creates a new GenreHandler.
func NewGenreHandler(store store.GenreStore) *GenreHandler {
	return &GenreHandler{store: store}
}

// CreateGenre handles the creation of a new genre.
func (h *GenreHandler) CreateGenre(w http.ResponseWriter, r *http.Request) {
	newGenre := &models.Genre{}

	// Get the genre data from the request body
	err := json.NewDecoder(r.Body).Decode(newGenre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the genre data
	if newGenre.Name == "" {
		http.Error(w, "Invalid genre data", http.StatusBadRequest)
		return
	}

	// Create the genre in the database
	_, err = h.store.CreateGenre(newGenre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetGenre handles the retrieval of a genre by ID.
func (h *GenreHandler) GetGenre(w http.ResponseWriter, r *http.Request) {
	idPath := r.PathValue("id")
	id, err := strconv.Atoi(idPath)
	if err != nil {
		http.Error(w, "Invalid genre ID", http.StatusBadRequest)
		return
	}

	genre, err := h.store.GetGenre(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(genre)
}

// GetGenres handles the retrieval of all genres.
func (h *GenreHandler) GetGenres(w http.ResponseWriter, r *http.Request) {
	var body getGenresBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	genres := []*models.Genre{}
	if body.Count <= 0  && body.Offset < 0 {
		genres, err = h.store.GetAllGenres()
	} else {
		genres, err = h.store.GetGenres(body.Count, body.Offset)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(genres)
}

// UpdateGenre handles the update of an existing genre.
func (h *GenreHandler) UpdateGenre(w http.ResponseWriter, r *http.Request) {
	idPath := r.PathValue("id")
	id, err := strconv.Atoi(idPath)
	if err != nil {
		http.Error(w, "Invalid genre ID", http.StatusBadRequest)
		return
	}

	updatedGenre := &models.Genre{}
	updatedGenre.ID = id
	err = json.NewDecoder(r.Body).Decode(updatedGenre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the genre data
	if updatedGenre.Name == "" {
		http.Error(w, "Invalid genre data", http.StatusBadRequest)
		return
	}

	// Update the genre in the database
	err = h.store.UpdateGenre(updatedGenre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DeleteGenre handles the deletion of a genre by ID.
func (h *GenreHandler) DeleteGenre(w http.ResponseWriter, r *http.Request) {
	idPath := r.PathValue("id")
	id, err := strconv.Atoi(idPath)
	if err != nil {
		http.Error(w, "Invalid genre ID", http.StatusBadRequest)
		return
	}

	err = h.store.DeleteGenre(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

