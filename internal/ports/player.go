package ports

import "github.com/nicolasAguilar180193/go-L/internal/domain"

type PlayerService interface {
	Create(player domain.Player) (id any, err error)
}
