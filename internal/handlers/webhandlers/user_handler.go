package webhandlers

import (
	"github.com/qazpalm/gig-agg/internal/auth"
	"github.com/qazpalm/gig-agg/internal/session"
	"github.com/qazpalm/gig-agg/internal/store"
)

// UserHandler is a struct that handles user-related HTTP requests.
type UserHandler struct {
	userStore     store.UserStore
	userAuthManager *auth.UserAuthManager
	sessionStore  *session.SessionStore
}


