package main

import (
	"log"

	"test/config"
	"test/internal/app"
	"test/internal/repository"
	"test/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize configuration and logger
	// config.InitConfig()
	// config.InitLogger()

	// Initialize database connection
	db := config.InitDB()

	// Initialize repository
	userRepository := repository.NewUserRepository(db)

	// Initialize service
	userService := service.NewUserService(userRepository)

	// Initialize Gin server
	router := gin.Default()

	// Register routes and handlers
	app.RegisterRoutes(router, userService)

	// Start the server
	log.Fatal(router.Run(":8080"))
}
