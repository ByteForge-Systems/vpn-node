package handlers

import (
	"net/http"

	"github.com/ByteForge-Systems/vpn-api/scripts"
	"github.com/gin-gonic/gin"
)

// Обработчик для клиентских эндпоинтов

func GetClientConfig(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user_id parameter"})
		return
	}
	vlessLink, err := scripts.GenerateVLESSLink(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"vlessLink": vlessLink})
}