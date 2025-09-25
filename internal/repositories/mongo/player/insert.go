package player

import (
	"context"
	"fmt"
	"time"

	"github.com/nicolasAguilar180193/go-L/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) Insert(player *domain.Player) (id any, err error) {
	player.ID = primitive.NewObjectID()
	now := time.Now()
	player.CreatedAt = &now
	player.UpdatedAt = &now

	collection := r.Client.Database("go-l").Collection("players")
	insertResult, err := collection.InsertOne(context.Background(), player)
	if err != nil {
		println("Error inserting document:", err.Error())
		return nil, fmt.Errorf("error inserting document: %w", err)
	}

	return insertResult.InsertedID, nil
}
