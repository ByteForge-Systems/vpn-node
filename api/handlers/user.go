package handlers

import (
	"net/http"

	"github.com/ByteForge-Systems/vpn-node/scripts"
	"github.com/gin-gonic/gin"
)
// Обработчик для эндпоинтов пользователей

func AddUser(c *gin.Context) {
    var request struct {
        UUID string `json:"uuid"`
    }

    if err := c.BindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }
    
    userID, err := scripts.GenerateUser(request.UUID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"id": userID})
}

func ListAllUsers(c *gin.Context) {
    users, err := scripts.ListUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"users": users})
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