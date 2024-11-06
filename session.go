package nest

import (
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

const sessionCookieName = "_nestGoSessionId"

// Session represents a user session
type Session struct {
	ID        string
	Values    sync.Map
	ExpiresAt time.Time
}

// SessionStore holds all sessions
var SessionStorage sync.Map

// NewSession creates a new session
func newSession() *Session {
	return &Session{
		ID:        uuid.NewString(), // Generate a unique session ID
		Values:    sync.Map{},
		ExpiresAt: time.Now().Add(24 * time.Hour), // Default expiration
	}
}

// SetExpiration sets the expiration time for the session
func (s *Session) SetExpiration(duration time.Duration) *Session {
	s.ExpiresAt = time.Now().Add(duration)
	return s
}

// Save writes the session data to the response cookie
func (s *Session) Save(c *Ctx, storageEngine ...func()) error {
	cookie := &http.Cookie{
		Name:     sessionCookieName,
		Value:    s.ID,
		Expires:  s.ExpiresAt,
		Path:     "/",
		HttpOnly: true, // Helps mitigate XSS attacks
		Secure:   true, // Use secure cookies in production (requires HTTPS)
	}
	http.SetCookie(c.Res(), cookie)

	// Store the session in the SessionStore
	SessionStorage.Store(s.ID, s)

	return nil
}

// GetValue retrieves a value from the session
func (s *Session) GetValue(key string) (interface{}, bool) {
	return s.Values.Load(key)
}

// SetValue sets a value in the session
func (s *Session) SetValue(key string, value interface{}) *Session {
	s.Values.Store(key, value)
	return s
}
