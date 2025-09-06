package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nicolasAguilar180193/go-L/cmd/api/core"
)

func main() {
	server := core.Server{
		GinEngine: gin.Default(),
	}

	server.GinEngine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Fatalln(server.GinEngine.Run(":8001"))
}
