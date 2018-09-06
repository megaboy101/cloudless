package main

import (
	"log"

	"github.com/megaboy101/cloudless/pkg/user"
)

func main() {
	// Load user service
	userService, err := user.New()

	if err != nil {
		panic(err)
	}

	_, err = userService.Register("jacobbleser@gmail.com", "megaboy101", "#megaboy101")

	if err != nil {
		log.Fatal(err)
	}
}
