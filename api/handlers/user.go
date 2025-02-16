package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ByteForge-Systems/vpn-node/scripts"
)

// Обработчик для эндпоинтов пользователей

func AddUser(c *gin.Context) {
	var user struct {
		ID string `json:"id"`
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := scripts.GenerateUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": userID})
}

func RemoveUser(c *gin.Context) {
	userID := c.Param("id")
	err := scripts.RemoveUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func GenerateVLESSLink(c *gin.Context) {
	userID := c.Param("id")
	link, err := scripts.GenerateVLESSLink(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"link": link})
}