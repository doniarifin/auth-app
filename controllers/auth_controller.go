package controllers

import (
	"auth-app/logics"
	"auth-app/models"
	"auth-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

// Register godoc
// @Description
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest false "Register Request"
// @Success 200 {object} RegisterResponse
// @Router /register [post]
func RegisterHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
			return
		}

		// check existing email
		existingUser, _ := logics.FindUserByEmail(db, req.Email)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing user"})
		// 	return
		// }
		if existingUser != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Email is already registered"})
			return
		}

		// Hash password
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		// Save user to DB
		newUser := &models.User{
			Email:    req.Email,
			Password: hashedPassword,
		}
		if err := models.CreateUser(db, newUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		// Generate JWT token
		token, err := utils.GenerateJWT(req.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusCreated, RegisterResponse{
			Message: "Registration successful",
			Token:   token,
		})
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

// Login godoc
// @Description
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest false "Login Request"
// @Success 200 {object} LoginResponse
// @Router /login [post]
func LoginHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest

		// check json
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON",
			})
			return
		}

		// check user by email
		user, err := logics.FindUserByEmail(db, req.Email)
		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User not found",
			})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Check password with hash
		if !utils.CheckPasswordHash(req.Password, user.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Wrong password",
			})
			return
		}

		// Generate JWT token
		token, err := utils.GenerateJWT(user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to generate token",
			})
			return
		}

		// Return token response
		c.JSON(http.StatusOK, LoginResponse{
			Message: "Success",
			Token:   token,
		})
	}
}
