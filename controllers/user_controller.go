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
