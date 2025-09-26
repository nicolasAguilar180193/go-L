package player

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicolasAguilar180193/go-L/cmd/api/core"
)

func (h *Handler) Get(c *gin.Context) {
	playerIdParam := c.Param("id")
	player, err := h.PlayerService.Get(playerIdParam)
	if err != nil {
		core.RespondError(c, err)
		return
	}
	c.JSON(http.StatusOK, player)
}
