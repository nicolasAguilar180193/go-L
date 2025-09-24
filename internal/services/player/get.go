package player

import (
	"github.com/nicolasAguilar180193/go-L/internal/domain"
)

func (s *Service) Get(id string) (player *domain.Player, err error) {
	player, err = s.Repository.Get(id)
	if err != nil {
		return nil, err
	}

	return player, nil
}
