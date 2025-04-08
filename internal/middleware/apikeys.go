package middleware

import (
	"net/http"
	"github.com/qazpalm/gig-agg/internal/apikeys"
)

// APIKeyMiddleware is a middleware that checks for a valid API key in the request header.
type APIKeyMiddleware struct {
	apiKeyManager *apikeys.ApiKeyManager
}

// NewAPIKeyMiddleware creates a new APIKeyMiddleware.
func NewAPIKeyMiddleware(apiKeyManager *apikeys.APIKeyManager) *APIKeyMiddleware {
	return &APIKeyMiddleware{apiKeyManager: apiKeyManager}
}

// ServeHTTP checks for a valid API key in the request header and calls the next handler if valid.
func (m *APIKeyMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		http.Error(w, "API key is required", http.StatusUnauthorized)
		return
	}

	if !m.apiKeyManager.IsValid(apiKey) {
		http.Error(w, "Invalid API key", http.StatusUnauthorized)
		return
	}

	next(w, r)
}