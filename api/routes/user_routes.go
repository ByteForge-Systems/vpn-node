package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ByteForge-Systems/vpn-node/api/handlers"
)

func SetupUserRoutes(router *gin.Engine) {
	user := router.Group("/api/key")
	{
		user.POST("/", handlers.AddUser)
		user.DELETE("/:id", handlers.RemoveUser)
		user.POST("/", handlers.ListAllUsers)
		user.GET("/:id/link", handlers.GenerateVLESSLink)
	}
}