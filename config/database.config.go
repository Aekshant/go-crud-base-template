package config

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:Root@123@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect Database---")
	}

	//Apply Migration
	err = db.AutoMigrate()
	if err != nil {
		log.Fatalf("Failed At Migration DB %v", err)
	}
	log.Println("Database Connected Successfully")
	return db
}
