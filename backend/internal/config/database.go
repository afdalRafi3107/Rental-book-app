package config

import (
	"log"
	"os"
	"rental-book/internal/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	EnvLoad()
	dsn := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB ERROR:", err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Book{}, &entity.Rental{})

	return db
}
