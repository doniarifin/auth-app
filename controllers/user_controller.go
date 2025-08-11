package controllers

import (
	"auth-app/logics"
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

// GetCurrentUser godoc
// @Security BearerAuth
// @Description Get current user by ID
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} EmailResponse
// @Router /api/v1/GetCurrentUser [get]
func GetCurrentUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetString("id")

		user, err := logics.FindUserByID(db, id)
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

// UpdateUser godoc
// @Security BearerAuth
// @Description Update a user's email
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param user body UserRequest true "User data"
// @Success 200 {object} UserReponse
// @Router /api/v1/UpdateUser/{id} [put]
func UpdateUser(db *gorm.DB, c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		currentUser, err := logics.FindUserByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
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

		update := models.UpdateUser(db, req, id)
		if update != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
	}
}

type DeleteRequest struct {
	ID string `json:"id"`
}

type DeleteReponse struct {
	Message string `json:"message"`
}

// @BasePath /api/v1

// DeleteUserByID godoc
// @Security BearerAuth
// @Description Delete a user given their ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} DeleteReponse
// @Router /api/v1/DeleteUser/{id} [delete]
func DeleteUser(db *gorm.DB, c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		getUser, err := logics.FindUserByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
			return
		}

		delete := models.DeleteUser(db, getUser.ID)
		if delete != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}
