package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ByteForge-Systems/vpn-api/api/routes"
)

// Точка входа

func main() {
	router := gin.Default()

	// Настройка маршрутов
	routes.SetupAdminRoutes(router)
	routes.SetupClientRoutes(router)

	// Запуск сервера
	router.Run(":8080")
}