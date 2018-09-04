package user

import (
	"errors"
)

// New creates a new user service
func New(r Repository) (*Service, error) {
	client, err := r.Connect()

	if err != nil {
		return nil, errors.New("[User Service] Injected repository could not connect")
	}

	return &Service{client}, nil
}
