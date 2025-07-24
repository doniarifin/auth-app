package controllers

import (
	"auth-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProfileHandler(db *gorm.DB) gin.HandlerFunc {
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
