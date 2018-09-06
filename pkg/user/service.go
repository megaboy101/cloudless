package user

type Servicer interface {
}

type Repository interface {
	Connect() (Repository, error)
	Register(email, username, password string) (*User, error)
}

type Service struct {
	Repository
}
