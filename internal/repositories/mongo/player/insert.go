package player

import (
	"context"
	"fmt"

	"github.com/nicolasAguilar180193/go-L/internal/domain"
)

func (r Repository) Insert(player domain.Player) (id any, err error) {

	collection := r.Client.Database("go-l").Collection("players")
	insertResult, err := collection.InsertOne(context.Background(), player)
	if err != nil {
		println("Error inserting document:", err.Error())
		return nil, fmt.Errorf("error inserting document: %w", err)
	}

	return insertResult.InsertedID, nil
}
