package main

import (
	"github.com/gin-gonic/gin"
	"github.com/megaboy101/cloudless/pkg/harbor"
)

var HarborService harbor.Servicer

func main() {
	r := gin.Default()

	hs, err := harbor.New()
	if err != nil {
		panic(err.Error())
	}

	HarborService = hs

	r.POST("/createNamespace", CreateNamespace)
	r.GET("/bootstrap", Bootstrap)

	r.Run(":8080")
}

type NamespaceCommand struct {
	IDToken string `json:"idtoken"`
}
