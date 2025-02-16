package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ByteForge-Systems/vpn-api/scripts"
)

// Обработчик для админских эндпоинтов

// Список всех пользователей
func ListUsers(c *gin.Context) {
	users, err := scripts.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Информация о конкретном пользователе
func GetUser(c *gin.Context) {
	userID := c.Param("id")
	users, err := scripts.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, user := range users {
		if user.ID == userID {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// Добавление нового пользователя
func AddUser(c *gin.Context) {
	userID, err := scripts.GenerateUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": userID})
}

// Удаление пользователя
func DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	err := scripts.RemoveUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

// Получение статуса Xray
func GetXrayStatus(c *gin.Context) {
	status, err := scripts.GetXrayStatus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": status})
}

// Перезапуск Xray
func RestartXray(c *gin.Context) {
	err := scripts.RestartXray()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Xray restarted"})
}

// Получение статистики сервера, когда-нибудь тут что-то начнет работать.
func GetServerStats(c *gin.Context) {
	metrics, err := scripts.GetServerMetrics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"metrics": metrics})
}