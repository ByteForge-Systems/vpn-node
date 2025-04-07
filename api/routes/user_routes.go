package routes

import (
	"github.com/ByteForge-Systems/vpn-node/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	user := router.Group("/api/key")
	{
		user.POST("/", handlers.AddUser)
		user.DELETE("/:id", handlers.RemoveUser)
		user.GET("/", handlers.ListAllUsers)
		user.GET("/:id/link", handlers.GenerateVLESSLink)
	}
}
