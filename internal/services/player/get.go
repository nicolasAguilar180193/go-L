package player

import (
	"fmt"

	"github.com/nicolasAguilar180193/go-L/internal/domain"
)

func (s *Service) Get(id string) (player *domain.Player, err error) {
	if id == "" {
		return nil, fmt.Errorf("player ID is required")
	}
	player, err = s.Repository.Get(id)
	if err != nil {
		return nil, fmt.Errorf("error getting player: %w", err)
	}

	return player, nil
}
