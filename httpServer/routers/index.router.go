package routers

import (
	"goGinTemplate/dependencyinjector"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, controllers dependencyinjector.Controllers) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	api := router.Group("/api")
	UserRoutes(api.Group("/users"), controllers.UserController)
}
