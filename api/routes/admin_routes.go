package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ByteForge-Systems/vpn-api/api/handlers"
)

// Регистрация админских эндпоинтов

func SetupAdminRoutes(router *gin.Engine) {
	admin := router.Group("/admin")
	{
		admin.GET("/users", handlers.ListUsers) //+
		admin.GET("/users/:id", handlers.GetUser) // +
		admin.POST("/users", handlers.AddUser) // +
		admin.DELETE("/users/:id", handlers.DeleteUser) // + 

		admin.GET("/xray/status", handlers.GetXrayStatus) // +
		admin.POST("/xray/restart", handlers.RestartXray) // +
		
		// добавить действия с серверами
		// новый сервер
		// удалить сервер
		// список серверов 

	}
}