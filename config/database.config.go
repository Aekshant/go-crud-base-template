package config

import (
	"os"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect Database---")
	}

	//Apply Migration
	// err = db.AutoMigrate(&model.User{})
	// if err != nil {
	// 	log.Fatalf("Failed At Migration DB %v", err)
	// }
	log.Println("Database Connected Successfully")
	return db
}
