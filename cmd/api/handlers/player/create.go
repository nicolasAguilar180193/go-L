package player

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicolasAguilar180193/go-L/internal/domain"
)

// Objetivo de los handlers:
// 1. recibir la request y traducirla
// 2. procesarla(consumir un servicio)
// 3. traducir la response

func (h *Handler) Create(c *gin.Context) {
	var playerCreateParams domain.Player
	if err := c.ShouldBindJSON(&playerCreateParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ===========================================

	insertResultID, err := h.PlayerService.Create(&playerCreateParams)
	if err != nil {
		var appErr domain.AppError
		if errors.As(err, &appErr) {
			if appErr.Code == domain.ErrCodeDuplicateKey {
				c.JSON(http.StatusConflict, appErr)
				return
			}
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oops! something went wrong :("})
		return
	}

	// ===========================================

	c.JSON(http.StatusCreated, gin.H{"insertedID": insertResultID})
}
