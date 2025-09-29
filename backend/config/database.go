package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	var (
		DBHost = os.Getenv("DB_HOST")
		DBPort = os.Getenv("DB_PORT")
		DBName = os.Getenv("DB_NAME")
		DBUser = os.Getenv("DB_USER")
		DBPass = os.Getenv("DB_PASSWORD")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", DBHost, DBUser, DBPass, DBName, DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	err = db.AutoMigrate(
	)

	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	return db
}