package main

import (
	"goGinTemplate/dependencyinjector"
	"goGinTemplate/httpServer/routers"

	"github.com/gin-gonic/gin"
	env "github.com/joho/godotenv"
)

func init() {
	env.Load()
}

func main() {
	// Initialize all dependencies
	controllers := dependencyinjector.InitAllDependencyInjectors()

	// Create a new Gin router
	router := gin.Default()

	// Initialize routes with injected dependencies
	routers.InitRoutes(router, controllers)

	// Start the server
	router.Run(":3000") // Use a configurable port if needed
}
