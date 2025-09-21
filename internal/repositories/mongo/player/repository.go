package player

import "go.mongodb.org/mongo-driver/v2/mongo"

type Repository struct {
	Client *mongo.Client
}
