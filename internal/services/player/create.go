package player

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/nicolasAguilar180193/go-L/internal/domain"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (s Service) Create(player domain.Player) (id any, err error) {
	player.CreatedAt = time.Now().UTC()

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
	insertResult, err := collection.InsertOne(ctx, player)
	if err != nil {
		log.Fatal(err)
	}

	return insertResult.InsertedID, nil
}
