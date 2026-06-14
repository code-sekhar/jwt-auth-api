package controllers

import (
	"jwt-auth-api/config"
	"jwt-auth-api/models"
	"jwt-auth-api/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	RoleID   uint   `json:"role_id"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	//var user models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(400, gin.H{
			"error": "Email already exists",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		RoleID:   input.RoleID,
	}
	config.DB.Create(&user)
	c.JSON(201, gin.H{
		"message": "User registered successfully",
	})

}
func Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	token, _ := utils.GenerateToken(
		user.ID.String(),
		user.RoleID,
	)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to generate token",
		})
		return
	}
	c.JSON(200, gin.H{
		"token": token,
	})

}
func Profile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	var user models.User
	if err := config.DB.Preload("Role").Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, gin.H{
		"id":      user.ID,
		"name":    user.Name,
		"email":   user.Email,
		"role":    user.Role.Name,
		"role_id": user.RoleID,
	})
}
func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Admin Access Granted",
	})
}
