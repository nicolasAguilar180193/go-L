package ports

import "github.com/nicolasAguilar180193/go-L/pkg/domain"

type PlayerService interface {
	Get(id string) (player *domain.Player, err error)
	Create(player *domain.Player) (id any, err error)
	Delete(id string) (err error)
}

type PlayerRepository interface {
	Get(id string) (player *domain.Player, err error)
	Insert(player *domain.Player) (id any, err error)
	Delete(id string) (err error)
}
