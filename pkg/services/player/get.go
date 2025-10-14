package player

import (
	"errors"
	"fmt"
	"log"

	"github.com/nicolasAguilar180193/go-L/pkg/domain"
)

func (s *Service) Get(id string) (player *domain.Player, err error) {

	if id == "" {
		return nil, fmt.Errorf("player ID is required")
	}

	player, err = s.Repository.Get(id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, domain.NewAppError(
				domain.ErrCodeNotFound,
				fmt.Sprintf("player with id '%s' not found", id))
		}
		if errors.Is(err, domain.ErrTimeout) {
			return nil, domain.NewAppError(
				domain.ErrCodeTimeout,
				"timeout error, try again later")
		}
		log.Println(err.Error())
		return nil, fmt.Errorf("unexpected error getting player: %w", err)
	}

	return player, nil
}
