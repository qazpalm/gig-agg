package middleware

import (
	"net/http"
	"github.com/qazpalm/gig-agg/internal/session"
)

type SessionMiddleware struct {
	sessionStore *session.SessionStore
}

func NewSessionMiddleware(sessionStore *session.SessionStore) *SessionMiddleware {
	return &SessionMiddleware{sessionStore: sessionStore}
}

func (sm *SessionMiddleware) ServeSessionProtected(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement session management logic
		next.ServeHTTP(w, r)
	}
}