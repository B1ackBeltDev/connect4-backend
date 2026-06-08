package auth

import "github.com/B1ackBeltDev/connect4-backend/internals/sessions"

type authRepo struct {
	sessions *sessions.Sessions
}

func (r *authRepo) checkIfUserExists(payload *registerPayload) (bool, error) {
	// TODO
	return false, nil
}

func (r *authRepo) createUser(payload *registerPayload) error {
	// TODO
	return nil
}

func (r *authRepo) getUserID(userName string) (int, error) {
	// TODO
	return 0, nil
}
