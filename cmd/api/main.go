package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nicolasAguilar180193/go-L/cmd/api/handlers/player"
	"github.com/nicolasAguilar180193/go-L/internal/repositories/mongo"
	playerMongo "github.com/nicolasAguilar180193/go-L/internal/repositories/mongo/player"
	playerService "github.com/nicolasAguilar180193/go-L/internal/services/player"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ginEngine := gin.Default()

	mongoClient, err := mongo.ConnectClient(os.Getenv("MONGODB_URI"))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	playerRepo := &playerMongo.Repository{
		Client: mongoClient,
	}

	playerSrv := &playerService.Service{
		Repository: playerRepo,
	}

	playerHandler := &player.Handler{
		PlayerService: playerSrv,
	}

	ginEngine.POST("/players", playerHandler.Create)

	ginEngine.GET("/players/:id", playerHandler.Get)

	ginEngine.DELETE("/players/:id", playerHandler.Delete)

	log.Fatalln(ginEngine.Run(":8001"))
}
