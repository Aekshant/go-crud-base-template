package main

import (
	userService "goGinTemplate/app/usecases/usersUsecases"
	"goGinTemplate/config"
	userController "goGinTemplate/httpServer/handler"
	mysqlrepo "goGinTemplate/infra/mysqlRepo"

	env "github.com/joho/godotenv"
)

func init() {
	env.Load()
}

func main() {
	db := config.InitDB()
	userRepositoryImpl := mysqlrepo.UserRepositoryInit(db)
	userServiceImpl := userService.UserServiceInit(userRepositoryImpl)
	userControllerImpl := userController.UserControllerInit(userServiceImpl)
	userControllerImpl.GetAllUserData()
}
