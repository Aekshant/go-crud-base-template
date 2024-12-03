package dependencyinjector

import (
	"goGinTemplate/config"
	userController "goGinTemplate/httpServer/handler"
)

type Controllers struct {
	UserController userController.UserController
}

func InitAllDependencyInjectors() Controllers {
	gormDB := config.InitDB()
	userInjector := InitUserInjector(gormDB)
	return Controllers{
		UserController: userInjector.UserCtrl,
	}
}
