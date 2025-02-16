package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ByteForge-Systems/vpn-api/api/handlers"
)

// Регистрация клиентских эндпоинтов

func SetupClientRoutes(router *gin.Engine) {
	client := router.Group("/client")
	{
		client.GET("/config", handlers.GetClientConfig)
	}
}