package controllers

import (
	"auth-app/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProfileHandler(db *sql.DB) gin.HandlerFunc {
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
