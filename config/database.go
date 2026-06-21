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
		&models.Permission{},
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

	permissions := []models.Permission{
		{Name: "create_user"},
		{Name: "view_user"},
		{Name: "update_user"},
		{Name: "delete_user"},
		{Name: "create_product"},
		{Name: "view_product"},
		{Name: "update_product"},
		{Name: "delete_product"},
	}
	// Assign permissions to roles
	var superAdminRole models.Role
	db.Where("name = ?", "Super Admin").First(&superAdminRole)
	var allPermissions []models.Permission
	db.Find(&allPermissions)
	db.Model(&superAdminRole).Association("Permissions").Replace(allPermissions)
	//Admin role permissions
	var adminRole models.Role
	db.Where("name = ?", "Admin").First(&adminRole)
	var adminPermissions []models.Permission
	db.Where("name IN ?", []string{"create_user", "view_user", "update_product"}).Find(&adminPermissions)
	db.Model(&adminRole).Association("Permissions").Replace(adminPermissions)
	//Manager role permissions
	var managerRole models.Role
	db.Where("name = ?", "Manager").First(&managerRole)
	var managerPermissions []models.Permission
	db.Where("name IN ?", []string{"view_product"}).Find(&managerPermissions)
	db.Model(&managerRole).Association("Permissions").Replace(managerPermissions)
	//User role permissions
	var userRole models.Role
	db.Where("name = ?", "User").First(&userRole)
	var userPermissions []models.Permission
	db.Where("name IN ?", []string{"view_product"}).Find(&userPermissions)
	db.Model(&userRole).Association("Permissions").Replace(userPermissions)

	for _, role := range roles {
		db.Where("name = ?", role.Name).FirstOrCreate(&role)
	}
	for _, permission := range permissions {
		db.FirstOrCreate(&permission, models.Permission{Name: permission.Name})
	}

	DB = db
	fmt.Println("Database connected successfully")
}
