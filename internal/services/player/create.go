package player

import (
	"fmt"
	"log"
	"time"

	"github.com/nicolasAguilar180193/go-L/internal/domain"
)

// pasos:
// 1. Set createdAt
// 2. Save to repository
// 3. Return ID of created player

func (s Service) Create(player domain.Player) (id any, err error) {
	now := time.Now().UTC()
	player.CreatedAt = &now

	// ========================================================

	insertedID, err := s.Repository.Insert(player)
	if err != nil {
		log.Println("Error saving player:", err.Error())
		return nil, fmt.Errorf("error saving player: %w", err)
	}

	// ========================================================

	return insertedID, nil
}
