package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ByteForge-Systems/vpn-api/scripts"
)

func RestartXray(c *gin.Context) {
	err := scripts.RestartXray()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Xray restarted"})
}

func GetXrayStatus(c *gin.Context) {
	status, err := scripts.GetXrayStatus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": status})
}