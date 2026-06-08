package auth

import (
	"os"
	"time"
)

type authService struct {
	repo authRepo
}

func (s *authService) register(payload *registerPayload) (string, error) {
	// 1. Check if user already exists
	exists, err := s.repo.checkIfUserExists(payload)
	if err != nil {
		return "", err
	}

	if !exists {
		// 2. Write user to database
		err := s.repo.createUser(payload)
		if err != nil {
			return "", err
		}
	}

	// 3. Get userID
	userID, err := s.repo.getUserID(payload.UserName)
	if err != nil {
		return "", err
	}

	// 3. Check weather session exists for that user
	session, err := s.repo.sessions.GetSessionFromUser(userID)
	if err != nil {
		return "", err
	}

	// 4. If no session is found or session has already expired create new session
	if err == os.ErrNotExist || session.ExpiresAt.Before(time.Now()) {
		if err := s.repo.sessions.CreateSession(userID); err != nil {
			return "", err
		}

	}

	// 5. Retrieve session and update expiration date

	// 6. Return session key

	return "123", nil
}
