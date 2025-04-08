package auth

import (
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/qazpalm/gig-agg/internal/store"
	"github.com/qazpalm/gig-agg/internal/models"
	"github.com/qazpalm/gig-agg/internal/session"
)

type UserAuthManager struct {
	userStore store.UserStore
	sessionStore *session.SessionStore
}

func NewUserAuthManager(userStore store.UserStore) *UserAuthManager {
	return &UserAuthManager{userStore: userStore}
}

func (uam *UserAuthManager) AuthenticateUser(email, password string) (models.User, error) {
	user, err := uam.userStore.GetUserByEmail(email)
	if err != nil {
		return models.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return models.User{}, err
	}

	// Update the last login time
	user.LastLogin = time.Now()
	if err := uam.userStore.UpdateUser(user); err != nil {
		return models.User{}, err
	}

	// Create a session for the user
	sessionToken := session.GenerateSessionToken()
	expiresAt := time.Now().Add(24 * time.Hour) // Set session expiration to 24 hours
	uam.sessionStore.AddSession(sessionToken, user.ID, expiresAt)
	user.RememberedToken = sessionToken
	if err := uam.userStore.UpdateUser(user); err != nil {
		return models.User{}, err
	}

	return *user, nil
}

func (uam *UserAuthManager) RegisterUser(email, username, password string) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Email:    email,
		Username: username,
		PasswordHash: string(hashedPassword),
	}

	userID, err := uam.userStore.CreateUser(&user)
	if err != nil {
		return models.User{}, err
	}

	user.ID = int(userID)
	return user, nil
}

