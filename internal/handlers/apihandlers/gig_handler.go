package apihandlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/qazpalm/gig-agg/internal/models"
    "github.com/qazpalm/gig-agg/internal/store"
)

// GigHandler handles requests related to gigs.
type GigHandler struct {
    store store.GigStore
}

// NewGigHandler creates a new GigHandler.
func NewGigHandler(store store.GigStore) *GigHandler {
    return &GigHandler{store: store}
}

// CreateGig handles the creation of a new gig.
func (h *GigHandler) CreateGig(w http.ResponseWriter, r *http.Request) {
    newGig := &models.Gig{}

    // Decode the gig data from the request body
    err := json.NewDecoder(r.Body).Decode(newGig)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Validate the gig data
    if newGig.Name == "" || newGig.VenueID == 0 || newGig.DateTime.IsZero() {
        http.Error(w, "Invalid gig data", http.StatusBadRequest)
        return
    }

    // Create the gig in the database
    _, err = h.store.CreateGig(newGig)
    if err != nil {
        http.Error(w, "Failed to create gig", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

// GetGig handles the retrieval of a gig by ID.
func (h *GigHandler) GetGig(w http.ResponseWriter, r *http.Request) {
    idPath := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idPath)
    if err != nil {
        http.Error(w, "Invalid gig ID", http.StatusBadRequest)
        return
    }

    gig, err := h.store.GetGig(id)
    if err != nil {
        http.Error(w, "Failed to retrieve gig", http.StatusInternalServerError)
        return
    }

    if gig == nil {
        http.Error(w, "Gig not found", http.StatusNotFound)
        return
    }

    // Return the gig data as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(gig)
}

// GetGigs handles the retrieval of all gigs with pagination.
func (h *GigHandler) GetGigs(w http.ResponseWriter, r *http.Request) {
    countStr := r.URL.Query().Get("count")
    offsetStr := r.URL.Query().Get("offset")

    count, err := strconv.Atoi(countStr)
    if err != nil || count <= 0 {
        count = 10 // Default to 10 if not provided or invalid
    }

    offset, err := strconv.Atoi(offsetStr)
    if err != nil || offset < 0 {
        offset = 0 // Default to 0 if not provided or invalid
    }

    gigs := []*models.Gig{}
    if count <= 0 && offset < 0 {
        gigs, err = h.store.GetAllGigs()
    } else {
        gigs, err = h.store.GetGigs(count, offset)
    }
    if err != nil {
        http.Error(w, "Failed to retrieve gigs", http.StatusInternalServerError)
        return
    }

    // Return the gigs data as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(gigs)
}

// UpdateGig handles the update of an existing gig.
func (h *GigHandler) UpdateGig(w http.ResponseWriter, r *http.Request) {
    idPath := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idPath)
    if err != nil {
        http.Error(w, "Invalid gig ID", http.StatusBadRequest)
        return
    }

    updatedGig := &models.Gig{}
    updatedGig.ID = id

    // Decode the updated gig data from the request body
    err = json.NewDecoder(r.Body).Decode(updatedGig)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Validate the updated gig data
    if updatedGig.Name == "" || updatedGig.VenueID == 0 || updatedGig.DateTime.IsZero() {
        http.Error(w, "Invalid gig data", http.StatusBadRequest)
        return
    }

    // Update the gig in the database
    err = h.store.UpdateGig(updatedGig)
    if err != nil {
        http.Error(w, "Failed to update gig", http.StatusInternalServerError)
        return
    }
}

// DeleteGig handles the deletion of a gig by ID.
func (h *GigHandler) DeleteGig(w http.ResponseWriter, r *http.Request) {
    idPath := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idPath)
    if err != nil {
        http.Error(w, "Invalid gig ID", http.StatusBadRequest)
        return
    }

    err = h.store.DeleteGig(id)
    if err != nil {
        http.Error(w, "Failed to delete gig", http.StatusInternalServerError)
        return
    }
}