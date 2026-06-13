package routes

import (
	"jwt-auth-api/controllers"
	"jwt-auth-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	protected := router.Group("/")

	protected.Use(
		middleware.AuthMiddleware(),
	)
	protected.GET("/profile", controllers.Profile)
}
