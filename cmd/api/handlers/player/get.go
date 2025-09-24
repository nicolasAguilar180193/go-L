package player

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Get(c *gin.Context) {
	playerIdParam := c.Param("id")
	player, err := h.PlayerService.Get(playerIdParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, player)
}
