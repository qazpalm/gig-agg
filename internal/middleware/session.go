package middleware

import (
	"net/http"
	"github.com/qazpalm/gig-agg/internal/session"
	"github.com/qazpalm/gig-agg/internal/store"
)

type SessionMiddleware struct {
	sessionStore *session.SessionStore
	userStore    *store.UserStore
}

func NewSessionMiddleware(sessionStore *session.SessionStore) *SessionMiddleware {
	return &SessionMiddleware{sessionStore: sessionStore}
}

func (sm *SessionMiddleware) ServeSessionProtected(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionToken := r.Header.Get("session_token")
		sessionData, exists := sm.sessionStore.GetSession(sessionToken)
		if !exists {
			rememberToken := r.Header.Get("remember_token")
			// TODO: Check if rememberToken is valid
			user, err := sm.userStore.GetUserByRememberedToken(rememberToken)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			newSessionToken := "some_new_session_token" // Generate a new session token
			// Create a new session for the user
			sm.sessionStore.AddSession(
				newSessionToken,
				user.ID,
				time.Now().Add(24*time.Hour), // Set session expiration to 24 hours
			)

			// Set the new session token in the response header
			w.Header().Set("session_token", newSessionToken)
		}
		next.ServeHTTP(w, r)
	}
}