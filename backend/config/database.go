package config

import (
	"fmt"
	"log"
	"training-plan-api/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDB(config *Config) *gorm.DB {
	dsn := config.DBUser + ":" + config.DBPass + "@tcp(" + config.DBHost + ":" + config.DBPort + ")/" + config.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	//  Auto Migration
	err = db.AutoMigrate(&model.Department{},&model.Course{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	fmt.Println("Connected and migrated successfully to the database")
	return db
}