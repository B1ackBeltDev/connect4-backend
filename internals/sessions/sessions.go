package sessions

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/B1ackBeltDev/connect4-backend/internals/utils"
)

type Sessions struct {
	sessions SessionsStorage
	users    SessionsStorage
}

type Session struct {
	SessionToken string    `json:"session_token"`
	UserID       int       `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func (s *Session) UpdateExpirationDate(duration time.Duration) {
	s.ExpiresAt = time.Now().Add(duration)
}

func (s *Sessions) CreateSession(userID int) error {
	sessionToken, err := utils.GenRandomToken(16)
	if err != nil {
		return err
	}

	session := Session{
		SessionToken: sessionToken,
		UserID:       userID,
		CreatedAt:    time.Now(),
		ExpiresAt:    time.Now().Add(time.Hour),
	}

	// Marshall session
	res, err := json.Marshal(&session)
	if err != nil {
		return err
	}

	marshaled := string(res)

	// For now do them sequeitally and assume that both of them dont fail
	err1 := s.sessions.Store(sessionToken, marshaled)
	if err1 != nil {
		return err1
	}

	err2 := s.users.Store(strconv.Itoa(userID), sessionToken)
	if err2 != nil {
		return err2
	}

	return nil
}

func (s *Sessions) GetSessionFromUser(userID int) (Session, error) {
	sessionToken, err := s.users.Get(strconv.Itoa(userID))
	if err != nil {
		return Session{}, err
	}

	// Unmarshall sessionToken
	var session Session
	if err := json.Unmarshal([]byte(sessionToken), &session); err != nil {
		return Session{}, err
	}

	return session, nil
}
