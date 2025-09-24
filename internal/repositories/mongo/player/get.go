package player

import (
	"context"
	"fmt"

	"github.com/nicolasAguilar180193/go-L/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Get(id string) (player *domain.Player, err error) {
	playerID, err := primitive.ObjectIDFromHex(id)
	// playerID, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	collection := r.Client.Database("go-l").Collection("players")
	err = collection.FindOne(context.Background(), bson.M{"_id": playerID}).Decode(&player)
	if err != nil {
		return nil, err
	}
	// ========================================================
	return player, nil
}
