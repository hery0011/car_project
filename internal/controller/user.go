package controller

import (
	"car_project/internal/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatUser handles the creation of a new user within a transactional context.
// It processes the incoming HTTP request, validates the input, and creates a user record in the database.
// If any error occurs during the process, the transaction is rolled back and an appropriate error response is returned.
// This handler expects a JSON payload representing the user to be created.
//
//	@Summary		Create a new user
//	@Description	Creates a new user in the system within a transactional context.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		entities.LoginStruct	true	"User creation payload"
//	@Success		200		{object}	map[string]interface{}	"User created successfully"
//	@Failure		400		{object}	map[string]interface{}	"Invalid request data"
//	@Failure		500		{object}	map[string]interface{}	"Internal server error"
//	@Router			/user/creatUser [post]
func (h *livraisonHandler) CreatUser(c *gin.Context) {
	var user entities.LoginStruct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid request data",
		})
		return
	}

	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Database connection not initialized",
		})
		return
	}
	tx := h.db.Begin()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to creat user",
		})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to commit transaction",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User created successfully",
		"data":    user,
	})
}
