package routes

import (
	"auth-app/controllers"
	"auth-app/middleware"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	// Public Routes
	r.POST("/login", controllers.LoginHandler(db))
	r.POST("/register", controllers.RegisterHandler(db))

	// Protected Routes - Require JWT
	protected := r.Group("/api")
	protected.Use(middleware.JWTMiddleware())

	protected.GET("/me", controllers.ProfileHandler(db))
}
