package main

import (
	"github.com/megaboy101/cloudless-api/pkg/user"
)

func main() {
	// Load user service
	userService, err := user.New()

	if err != nil {
		panic(err)
	}

	userService.RegisterUser("jacobbleser@gmail.com", "megaboy101", "#megaboy101")
}
