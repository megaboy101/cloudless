package user

import (
	"errors"
)

// New creates a new user service
func New() (*Service, error) {
	firebaseRepo := &FirebaseRepository{}

	client, err := firebaseRepo.Connect()

	if err != nil {
		return nil, errors.New("[User Service] Injected repository could not connect")
	}

	return &Service{client}, nil
}
