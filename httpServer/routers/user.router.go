package routers

import (
	userController "goGinTemplate/httpServer/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, userCtrl userController.UserController) {
	router.GET("/", userCtrl.GetAllUserData)
}
