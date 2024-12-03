package main

import (
	"goGinTemplate/config"
	diInit "goGinTemplate/dependencyInjector"

	env "github.com/joho/godotenv"
)

func init() {
	env.Load()
}

func main() {
	db := config.InitDB()
	init := diInit.Init(db)
	init.UserCtrl.GetAllUserData()
}
