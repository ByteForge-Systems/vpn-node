package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ByteForge-Systems/vpn-node/api/handlers"
)

func SetupManagementRoutes(router *gin.Engine) {
	management := router.Group("/api/management")
	{
		management.POST("/start", handlers.StartXray)
		management.POST("/stop", handlers.StopXray)
		management.POST("/restart", handlers.RestartXray)
		management.GET("/status", handlers.GetXrayStatus)
	}
}