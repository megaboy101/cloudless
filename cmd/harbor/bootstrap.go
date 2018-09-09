package main

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func Bootstrap(c *gin.Context) {
	if HarborService == nil {
		panic(errors.New("Could not bootstrap, Harbor is not initialized"))
	}

	if err := HarborService.InstallTiller(); err != nil {
		panic(err.Error())
	}

	c.String(200, "Tiller successfully installed!")
}
