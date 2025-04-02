package handlers

import (
	"net/http"
)

// HomeHandler handles the home page request.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("<h1>Welcome to Gig Aggregator</h1>"))
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

