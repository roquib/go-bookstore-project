package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gobookstore?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	db = d
	// Use db for CRUD operations
	log.Println("Database connected successfully!")
}

func GetDB() *gorm.DB {
	return db
}
