package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:root@123@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Apply migrations
	// err = db.AutoMigrate(&dao.User{})
	// if err != nil {
	// 	log.Fatalf("failed to migrate database: %v", err)
	// }

	// log.Println("Database connected and migrated successfully.")

	return db
}
