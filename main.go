package main

import (
	"goGinTemplate/config"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	config.InitLog()
	config.InitDB()
}

func main() {
	// port := os.Getenv("PORT")

}
