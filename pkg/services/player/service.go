package player

import "github.com/nicolasAguilar180193/go-L/pkg/ports"

// Make sure Service implements the PlayerService interface
// at compile time.
var _ ports.PlayerService = &Service{}

type Service struct {
	Repository ports.PlayerRepository
}
