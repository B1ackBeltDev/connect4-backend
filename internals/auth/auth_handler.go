package auth

import (
	"encoding/json"
	"errors"
	"net/http"
)

type AuthHandler struct {
	service authService
}

type registerPayload struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) registerHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1024*1024)

	// Decode payload
	var payload registerPayload

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&payload); err != nil {
		var maxBytesErr *http.MaxBytesError
		if errors.As(err, &maxBytesErr) {
			http.Error(w, "Requested Body is too big (max 1MB)", http.StatusRequestEntityTooLarge)
			return
		}

		http.Error(w, "Failed to decode JSON Body", http.StatusBadRequest)
		return
	}

	// Validate if fields are not null
	switch {
	case len(payload.Email) == 0:
		http.Error(w, "Email field must not be empty", http.StatusBadRequest)
		return
	case len(payload.Password) == 0:
		http.Error(w, "Password field must not be empty", http.StatusBadRequest)
		return
	case len(payload.UserName) == 0:
		http.Error(w, "User name field must not be empty", http.StatusBadRequest)
		return
	}


	
}

func (h *AuthHandler) loginHandler(w http.ResponseWriter, r *http.Request) {

}
