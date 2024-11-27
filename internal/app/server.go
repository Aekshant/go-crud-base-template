package app

import (
	"test/internal/controller"
	"test/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, userService service.UserService) {
	userController := controller.NewUserController(userService)
	userController.RegisterRoutes(router)
}
