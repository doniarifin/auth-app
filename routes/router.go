package routes

import (
	"auth-app/controllers"
	"auth-app/docs"
	"auth-app/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Terapkan middleware CORS ke router utama (r)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	docs.SwaggerInfo.BasePath = ""

	// Public Routes
	r.POST("/login", controllers.LoginHandler(db))
	r.POST("/register", controllers.RegisterHandler(db))

	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Protected Routes - Require JWT
	protected := r.Group("/api/v1")
	protected.Use(middleware.JWTMiddleware())
	protected.GET("/GetCurrentUser", controllers.GetCurrentUser(db))
}
