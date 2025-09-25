package ports

import "github.com/nicolasAguilar180193/go-L/internal/domain"

type PlayerService interface {
	Create(player *domain.Player) (id any, err error)
	Get(id string) (player *domain.Player, err error)
	Delete(id string) (err error)
}

type PlayerRepository interface {
	Insert(player *domain.Player) (id any, err error)
	Get(id string) (player *domain.Player, err error)
	Delete(id string) (err error)
}
