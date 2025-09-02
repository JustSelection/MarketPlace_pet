package db

import (
	"MarketPlace_Pet/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5454 sslmode=disable"

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	// Автомиграции! Убрать!
	err = db.AutoMigrate(
		&models.User{},
		&models.UserCartItem{},
		&models.Order{},
		&models.OrderItem{},
		&models.Product{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Database connected successfully")
	return db, nil
}
