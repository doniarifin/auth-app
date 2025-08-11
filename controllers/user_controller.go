package controllers

import (
	"auth-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EmailRequest struct {
	Email string `json:"email" example:"user@example.com"`
}
type EmailResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// @BasePath /api/v1

// PingExample godoc
// @Summary
// @Schemes
// @Security BearerAuth
// @Description
// @Tags GetCurrentUser
// @Accept json
// @Produce json
// @Success 200 {object} EmailResponse
// @Router /api/v1/GetCurrentUser [get]
func GetCurrentUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.GetString("email")

		user, err := models.FindUserByEmail(db, email)
		if err != nil || user == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":    user.ID,
			"email": user.Email,
		})
	}
}

type UserRequest struct {
	Email string `json:"email"`
}

type UserReponse struct {
	Message string `json:"message"`
}

// @BasePath /api/v1

// PingExample godoc
// @Summary
// @Schemes
// @Security BearerAuth
// @Description
// @Tags UpdateUser
// @Accept json
// @Produce json
// @Param request body UserRequest true "UserRequest"
// @Success 200 {object} UserReponse
// @Router /api/v1/UpdateUser [put]
func UpdateUser(db *gorm.DB, c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.GetString("email")

		currentUser, err := models.FindUserByEmail(db, email)
		if err != nil {
			return
		}

		var payload UserRequest
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		req := &models.User{
			ID:       currentUser.ID,
			Password: currentUser.Password,
			Email:    payload.Email,
		}

		update := models.UpdateUser(db, req)
		if update != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
	}
}
