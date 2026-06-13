package config

import (
	"fmt"
	"jwt-auth-api/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s  password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: ", err)
	}
	err = db.AutoMigrate(
		&models.Role{},
		&models.User{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: ", err)
	}
	roles := []models.Role{
		{Name: "Super Admin"},
		{Name: "Admin"},
		{Name: "Manager"},
		{Name: "User"},
	}
	for _, role := range roles {
		db.Where("name = ?", role.Name).FirstOrCreate(&role)
	}

	DB = db
	fmt.Println("Database connected successfully")
}
