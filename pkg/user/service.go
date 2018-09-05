package user

type Servicer interface {
}

type Repository interface {
	Connect() (Repository, error)
	RegisterUser(email, username, password string) (*User, error)
}

type Service struct {
	Repository
}
