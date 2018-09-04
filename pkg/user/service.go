package user

import (
	"fmt"
)

type Servicer interface {
	New() (*User, error)
	Hello()
}

type Repository interface {
	Connect() (Repository, error)
}

type Service struct {
	repository Repository
}

func (s *Service) Hello() {
	fmt.Println("Hello there, I am a new user!")
}
