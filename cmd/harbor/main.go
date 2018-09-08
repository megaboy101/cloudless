package main

import (
	"github.com/gin-gonic/gin"
	"github.com/megaboy101/cloudless/pkg/harbor"
)

func main() {
	r := gin.Default()

	r.POST("/createNamespace", func(c *gin.Context) {
		var namespaceCommand NamespaceCommand

		err := c.BindJSON(&namespaceCommand)
		if err != nil {
			panic(err.Error())
		}

		hs, err := harbor.New()
		if err != nil {
			panic(err.Error())
		}

		err = hs.CreateNamespace(namespaceCommand.IDToken)
		if err != nil {
			panic(err.Error())
		}

		c.JSON(200, gin.H{"ok": true})
	})

	r.Run(":8080")
}

type NamespaceCommand struct {
	IDToken string `json:"idtoken"`
}