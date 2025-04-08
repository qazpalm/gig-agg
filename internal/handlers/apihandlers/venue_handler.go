package apihandlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/qazpalm/gig-agg/internal/models"
    "github.com/qazpalm/gig-agg/internal/store"
)

// VenueHandler handles requests related to venues.
type VenueHandler struct {
    store store.VenueStore
}

// NewVenueHandler creates a new VenueHandler.
func NewVenueHandler(store store.VenueStore) *VenueHandler {
    return &VenueHandler{store: store}
}

// CreateVenue handles the creation of a new venue.
func (h *VenueHandler) CreateVenue(w http.ResponseWriter, r *http.Request) {
    newVenue := &models.Venue{}

    // Decode the venue data from the request body
    err := json.NewDecoder(r.Body).Decode(newVenue)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Validate the venue data
    if newVenue.Name == "" || newVenue.Address == "" || newVenue.City == "" {
        http.Error(w, "Invalid venue data", http.StatusBadRequest)
        return
    }

    // Create the venue in the database
    _, err = h.store.CreateVenue(newVenue)
    if err != nil {
        http.Error(w, "Failed to create venue", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

// GetVenue handles the retrieval of a venue by ID.
func (h *VenueHandler) GetVenue(w http.ResponseWriter, r *http.Request) {
    idPath := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idPath)
    if err != nil {
        http.Error(w, "Invalid venue ID", http.StatusBadRequest)
        return
    }

    venue, err := h.store.GetVenue(id)
    if err != nil {
        http.Error(w, "Failed to retrieve venue", http.StatusInternalServerError)
        return
    }

    if venue == nil {
        http.Error(w, "Venue not found", http.StatusNotFound)
        return
    }

    // Return the venue data as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(venue)
}

// GetVenues handles the retrieval of all venues with pagination.
func (h *VenueHandler) GetVenues(w http.ResponseWriter, r *http.Request) {
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

    venues, err := h.store.GetVenues(count, offset)
    if err != nil {
        http.Error(w, "Failed to retrieve venues", http.StatusInternalServerError)
        return
    }

    // Return the venues data as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(venues)
}

// UpdateVenue handles the update of an existing venue.
func (h *VenueHandler) UpdateVenue(w http.ResponseWriter, r *http.Request) {
    idPath := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idPath)
    if err != nil {
        http.Error(w, "Invalid venue ID", http.StatusBadRequest)
        return
    }

    updatedVenue := &models.Venue{}
    updatedVenue.ID = id

    // Decode the updated venue data from the request body
    err = json.NewDecoder(r.Body).Decode(updatedVenue)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Validate the updated venue data
    if updatedVenue.Name == "" || updatedVenue.Address == "" || updatedVenue.City == "" {
        http.Error(w, "Invalid venue data", http.StatusBadRequest)
        return
    }

    // Update the venue in the database
    err = h.store.UpdateVenue(updatedVenue)
    if err != nil {
        http.Error(w, "Failed to update venue", http.StatusInternalServerError)
        return
    }
}

// DeleteVenue handles the deletion of a venue by ID.
func (h *VenueHandler) DeleteVenue(w http.ResponseWriter, r *http.Request) {
    idPath := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idPath)
    if err != nil {
        http.Error(w, "Invalid venue ID", http.StatusBadRequest)
        return
    }

    err = h.store.DeleteVenue(id)
    if err != nil {
        http.Error(w, "Failed to delete venue", http.StatusInternalServerError)
        return
    }
}