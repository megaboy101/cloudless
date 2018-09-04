package main

import (
	"github.com/megaboy101/cloudless-api/pkg/user"
)

func main() {
	// Load user service
	firebaseRepo := &user.FirebaseRepository{}

	userService, err := user.New(firebaseRepo)

	if err != nil {
		panic(err)
	}

	userService.Hello()
}
