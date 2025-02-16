package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ByteForge-Systems/vpn-node/api/routes"
	"github.com/ByteForge-Systems/vpn-node/utils"
)

func main() {
	// Загрузка переменных окружения
	utils.LoadEnv()

	// Инициализация Gin
	router := gin.Default()

	// Настройка маршрутов
	routes.SetupUserRoutes(router)
	routes.SetupManagementRoutes(router)

	// Запуск сервера
	router.Run(":8080")
}