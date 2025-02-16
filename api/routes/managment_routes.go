package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ByteForge-Systems/vpn-api/api/handlers"
)

func SetupManagementRoutes(router *gin.Engine) {
	management := router.Group("/api/management")
	{
		management.POST("/restart", handlers.RestartXray)
		management.GET("/status", handlers.GetXrayStatus)
	}
}