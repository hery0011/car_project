package controller

import (
	"car_project/internal/entities"
	"fmt"
	"net/http"
	"strconv"

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
//	@Router			/admin/user/creatUser [post]
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

// DeleteUser supprime un utilisateur existant par son ID.
//
// @Summary      Supprimer un utilisateur
// @Description  Supprime un utilisateur en fonction de son identifiant (ID).
// @Tags         users
// @Param        idUser   path      int  true  "ID de l'utilisateur"
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "Suppression réussie"
// @Failure      400  {object}  map[string]interface{}  "ID invalide"
// @Failure      200  {object}  map[string]interface{}  "Utilisateur introuvable"
// @Failure      500  {object}  map[string]interface{}  "Erreur interne du serveur"
// @Router       /admin/user/{idUser}/delete [delete]
func (h *livraisonHandler) DeleteUser(c *gin.Context) {
	var userEntite entities.LoginStruct
	var idUser = c.Param("idUser")
	id, err := strconv.Atoi(idUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ID invalide",
		})
		return
	}

	if err := h.db.First(&userEntite, id).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Utilisateur introuvable",
		})
		return
	}

	if err := h.db.Where("id = ?", id).Delete(&userEntite).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "suppression user avec succès",
		"data":    userEntite,
	})
	fmt.Println(idUser)
}

// UpdateUser godoc
// @Summary Met à jour un utilisateur
// @Description Met à jour les informations d'un utilisateur existant
// @Tags users
// @Accept json
// @Produce json
// @Param payload body entities.LoginStruct true "Utilisateur à mettre à jour"
// @Success 200 {object} map[string]interface{} "Utilisateur mis à jour avec succès"
// @Failure 400 {object} map[string]interface{} "Erreur lors de la mise à jour"
// @Router /admin/user/updateUser [put]
func (h *livraisonHandler) UpdateUser(c *gin.Context) {
	var payload entities.LoginStruct

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid request data",
		})
		return
	}

	if payload.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ID utilisateur manquant",
		})
		return
	}

	if err := h.db.Where("id = ?", payload.Id).Updates(payload).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Utilisateur mis à jour avec succès",
		"data":    payload,
	})
}
