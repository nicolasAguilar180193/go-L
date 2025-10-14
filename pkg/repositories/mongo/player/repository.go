package player

import (
	"github.com/nicolasAguilar180193/go-L/pkg/ports"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ ports.PlayerRepository = &Repository{}

type Repository struct {
	Client *mongo.Client
}
