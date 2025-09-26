package player

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/nicolasAguilar180193/go-L/internal/domain"
)

// pasos:
// 1. Set createdAt
// 2. Save to repository
// 3. Return ID of created player

func (s *Service) Create(player *domain.Player) (id any, err error) {
	now := time.Now().UTC()
	player.CreatedAt = &now

	// ========================================================

	insertedID, err := s.Repository.Insert(player)
	if err != nil {
		if errors.Is(err, domain.ErrDuplicateKey) {
			log.Println("Duplicate key error.")
			appErr := domain.AppError{
				Code: domain.ErrCodeDuplicateKey,
				Msg:  "A player with the same unique field already exists",
			}
			return nil, appErr
		}
		log.Println(err.Error())
		return nil, fmt.Errorf("error saving player: %w", err)
	}

	// ========================================================

	return insertedID, nil
}
