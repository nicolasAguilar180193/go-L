package player

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Delete(c *gin.Context) {
	playerIdParam := c.Param("id")
	err := h.PlayerService.Delete(playerIdParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "Player deleted successfully"})
}
