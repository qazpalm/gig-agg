package session

import (
	"sync"
	"time"
)

type SessionData struct {
	UserID    int
	ExpiresAt time.Time
}

type SessionStore struct {
	sessions map[string]SessionData
	mu       sync.RWMutex
}

func NewSessionStore() *SessionStore {
	return &SessionStore{
		sessions: make(map[string]SessionData),
	}
}

func (s *SessionStore) AddSession(token string, userID int, expiresAt time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sessions[token] = SessionData{
		UserID:    userID,
		ExpiresAt: expiresAt,
	}
}

func (s *SessionStore) GetSession(token string) (SessionData, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	session, exists := s.sessions[token]
	if !exists || session.ExpiresAt.Before(time.Now()) {
		return SessionData{}, false
	}
	return session, true
}

func (s *SessionStore) DeleteSession(token string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.sessions, token)
}

func (s *SessionStore) CleanupExpiredSessions() {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	now := time.Now()
	for token, session := range s.sessions {
		if session.ExpiresAt.Before(now) {
			delete(s.sessions, token)
		}
	}
}