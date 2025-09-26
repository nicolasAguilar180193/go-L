package player

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nicolasAguilar180193/go-L/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) Insert(player *domain.Player) (id any, err error) {
	playerID := primitive.NewObjectID()
	player.ID = playerID
	now := time.Now()
	player.CreatedAt = &now
	player.UpdatedAt = &now

	collection := r.Client.Database("go-l").Collection("players")
	insertResult, err := collection.InsertOne(context.Background(), player)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			log.Println("Duplicate key error:", err.Error())
			return nil, fmt.Errorf("%w error inserting player: %s",
				domain.ErrDuplicateKey, err.Error())
		}
		log.Println("Error inserting player:", err.Error())
		return nil, fmt.Errorf("error inserting player: %w", err)
	}

	log.Println("Inserted player with ID:", insertResult.InsertedID)

	return playerID.Hex(), nil
}
