package dependencyinjector

// import (
// 	service "goGinTemplate/app/usecases/usersUsecases"
// 	controller "goGinTemplate/httpServer/handler"
// 	repository "goGinTemplate/infra/mysqlRepo"

// 	"gorm.io/gorm"
// )

// type Initialization struct {
// 	userRepo repository.UserRepository
// 	userSvc  service.UserService
// 	UserCtrl controller.UserController
// }

// func Init(gormDB *gorm.DB) *Initialization {
// 	userRepository := repository.UserRepositoryInit(gormDB)
// 	userService := service.UserServiceImpl(userRepository)
// 	userController := controller.UserControllerImpl(userService)

// 	return (&Initialization{
// 		UserController: userController,
// 		UserService:    userService,
// 		UserRepository: userRepository,
// 	})
// }
