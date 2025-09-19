package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nicolasAguilar180193/go-L/cmd/api/handlers/player"
	playerSrv "github.com/nicolasAguilar180193/go-L/internal/services/player"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ginEngine := gin.Default()

	playerService := playerSrv.Service{}

	playerHandler := player.Handler{
		PlayerService: playerService,
	}

	ginEngine.POST("/players", playerHandler.Create)

	log.Fatalln(ginEngine.Run(":8001"))
}
