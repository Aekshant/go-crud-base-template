package dependencyinjector

import (
	userService "goGinTemplate/app/usecases/usersUsecases"
	userController "goGinTemplate/httpServer/handler"
	repository "goGinTemplate/infra/mysqlRepo"

	"gorm.io/gorm"
)

type Initialization struct {
	userRepo repository.UserRepository
	userSvc  userService.UserService
	UserCtrl userController.UserController
}

func InitUserInjector(gormDB *gorm.DB) *Initialization {
	userRepositoryImpl := repository.UserRepositoryInit(gormDB)
	userServiceImpl := userService.UserServiceInit(userRepositoryImpl)
	userControllerImpl := userController.UserControllerInit(userServiceImpl)

	return (&Initialization{
		userRepo: userRepositoryImpl,
		userSvc:  userServiceImpl,
		UserCtrl: userControllerImpl,
	})
}
