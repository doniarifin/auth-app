package routes

import (
	"auth-app/controllers"
	"auth-app/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Public Routes
	r.POST("/login", controllers.LoginHandler(db))
	r.POST("/register", controllers.RegisterHandler(db))

	// Protected Routes - Require JWT
	protected := r.Group("/api")
	protected.Use(middleware.JWTMiddleware())

	protected.GET("/me", controllers.ProfileHandler(db))
}
