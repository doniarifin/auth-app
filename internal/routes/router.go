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
	//auth service
	srv := service.NewAuthService(repo)
	h := handler.NewAuthHandler(srv)

	//user service
	userSrv := service.NewUserService(repo)
	userHandler := handler.NewUserHandler(userSrv)

	// Public Routes
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)

	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Protected Routes - Require JWT
	protected := r.Group("/api/v1")
	protected.Use(middleware.JWTMiddleware())
	protected.GET("/GetAllUsers", userHandler.GetAllUsers)
	protected.GET("/GetCurrentUser", userHandler.GetCurrentUser)
	protected.PUT("/Update/:id", userHandler.Update)
	protected.DELETE("/Delete/:id", middleware.Authorization("admin"), userHandler.Delete)
}
