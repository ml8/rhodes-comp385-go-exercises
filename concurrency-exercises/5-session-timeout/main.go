package main

import (
	"errors"
	"flag"
	"log"

	"time" // XXX rm
)

// SessionManager keeps track of all sessions from creation, updating
// to destroying.
// TODO: Make thread safe and run periodic cleanup in a background thread.
type SessionManager struct {
	sessions map[string]Session
}

// Session stores the session's data
type Session struct {
	Data     map[string]interface{}
	Accessed time.Time
}

func (s Session) Expired() bool {
	return time.Now().Sub(s.Accessed) > 5*time.Second
}

// NewSessionManager creates a new sessionManager
func NewSessionManager() *SessionManager {
	m := &SessionManager{
		sessions: make(map[string]Session),
	}

	// TODO: This should start the cleanup thread.

	return m
}

// CreateSession creates a new session and returns the sessionID
func (m *SessionManager) CreateSession() (string, error) {
	sessionID, err := MakeSessionID()
	if err != nil {
		return "", err
	}

	m.sessions[sessionID] = Session{
		Data:     make(map[string]interface{}),
		Accessed: time.Now(), // XXX rm
	}

	return sessionID, nil
}

// ErrSessionNotFound returned when sessionID not listed in
// SessionManager
var ErrSessionNotFound = errors.New("SessionID does not exists")

// GetSessionData returns data related to session if sessionID is
// found, errors otherwise
func (m *SessionManager) GetSessionData(sessionID string) (map[string]interface{}, error) {
	session, ok := m.sessions[sessionID]
	if !ok {
		return nil, ErrSessionNotFound
	}
	return session.Data, nil
}

// UpdateSessionData overwrites the old session data with the new one
func (m *SessionManager) UpdateSessionData(sessionID string, data map[string]interface{}) error {
	_, ok := m.sessions[sessionID]
	if !ok {
		return ErrSessionNotFound
	}

	// Hint: you should renew expiry of the session here
	m.sessions[sessionID] = Session{
		Data:     data,
		Accessed: time.Now(), // XXX rm
	}

	return nil
}

func main() {
	flag.Parse()

	// Create new sessionManager and new session
	m := NewSessionManager()
	sID, err := m.CreateSession()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Created new session with ID", sID)

	// Update session data
	data := make(map[string]interface{})
	data["website"] = "rhodes.edu"

	err = m.UpdateSessionData(sID, data)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Update session data, set website to rhodes.edu")

	// Retrieve data from manager again
	updatedData, err := m.GetSessionData(sID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Get session data:", updatedData)
}
