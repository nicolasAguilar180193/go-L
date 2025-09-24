package player

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Delete(id string) (err error) {

	playerID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid player ID: %w", err)
	}

	collection := r.Client.Database("go-l").Collection("players")
	deleteResult, err := collection.DeleteOne(context.Background(), bson.M{"_id": playerID})

	if err != nil {
		return fmt.Errorf("failed to delete player: %w", err)
	}

	if deleteResult.DeletedCount == 0 {
		return fmt.Errorf("no player found with the given ID")
	}

	return nil
}
