package player

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Get(c *gin.Context) {
	playerIdParam := c.Param("id")
	player, err := h.PlayerService.Get(playerIdParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}
	c.JSON(http.StatusOK, player)
}
