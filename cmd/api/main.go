package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	//"github.com/nicolasAguilar180193/go-L/cmd/api/core"
)

type Player struct {
	Name      string `json:"name" binding:"required"`
	Age       int    `json:"age" binding:"required,min=0"`
	CreatedAt time.Time
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ginEngine := gin.Default()

	ginEngine.POST("/players", func(c *gin.Context) {
		var newPlayer Player
		if err := c.ShouldBindJSON(&newPlayer); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		// Add time to
		newPlayer.CreatedAt = time.Now().UTC()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
		if err != nil {
			log.Fatal(err)
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}

		collection := client.Database("go-l").Collection("players")
		insertResult, err := collection.InsertOne(ctx, newPlayer)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(200, gin.H{"player": newPlayer, "insertedID": insertResult.InsertedID})
	})

	log.Fatalln(ginEngine.Run(":8001"))
}
