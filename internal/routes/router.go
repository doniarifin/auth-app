package routes

import (
	"auth-app/docs"
	"auth-app/internal/handler"
	"auth-app/internal/middleware"
	"auth-app/internal/repository"
	"auth-app/internal/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// c := gin.Context{}

	// apply middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	docs.SwaggerInfo.BasePath = ""

	repo := repository.NewUserRepository(db)
	srv := service.NewUserService(repo)
	h := handler.NewUserHandler(srv)

	// Public Routes
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)

	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Protected Routes - Require JWT
	protected := r.Group("/api/v1")
	protected.Use(middleware.JWTMiddleware())
	// protected.GET("/GetCurrentUser", handler.GetCurrentUser(db))
	// protected.PUT("/UpdateUser/:id", handler.UpdateUser(db, c))
	// protected.DELETE("/DeleteUser/:id", handler.DeleteUser(db, c))
}
