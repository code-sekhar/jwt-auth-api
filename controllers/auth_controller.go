package controllers

import (
	"jwt-auth-api/config"
	"jwt-auth-api/models"
	"jwt-auth-api/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var existingUser models.User
	if err := config.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(400, gin.H{
			"error": "Email already exists",
		})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
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
	c.ShouldBindJSON(&req)
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
	token, _ := utils.GenerateToken(user.ID.String())
	c.JSON(200, gin.H{
		"token": token,
	})

}
func Profile(c *gin.Context) {
	c.JSON(200, gin.H{"message": "This is a protected route|| Welcome to your profile!"})
}
