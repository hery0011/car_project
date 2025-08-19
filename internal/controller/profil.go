package controller

import (
	"car_project/internal/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetListProfil godoc
// @Summary      Récupérer la liste des profils
// @Description  Retourne la liste complète des profils disponibles
// @Tags         profil
// @Produce      json
// @Success      200 {object} map[string]interface{} "Liste des profils récupérée avec succès"
// @Failure      404 {object} map[string]interface{} "Aucun profil trouvé"
// @Failure      500 {object} map[string]interface{} "Erreur lors de la récupération des profils"
// @Router       /admin/profil/list [get]
func (h *livraisonHandler) GetListProfil(c *gin.Context) {
	var profils []entities.Profil

	// Récupération des profils
	if err := h.db.Find(&profils).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Erreur lors de la récupération des profils",
			"error":   err.Error(),
		})
		return
	}

	// Si aucun profil trouvé
	if len(profils) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Aucun profil trouvé",
			"data":    []entities.Profil{},
		})
		return
	}

	// Succès
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Liste des profils récupérée avec succès",
		"data":    profils,
	})
}

// AssignationProfilUser godoc
// @Summary      Assigner un profil à un utilisateur
// @Description  Permet d'associer un profil (client, commerçant, livreur, etc.) à un utilisateur existant
// @Tags         profil
// @Accept       json
// @Produce      json
// @Param        payload body entities.PayloadAssignProfil true "Données pour l’assignation du profil"
// @Success      200 {object} map[string]interface{} "Profil assigné avec succès"
// @Failure      400 {object} map[string]interface{} "Payload invalide"
// @Failure      500 {object} map[string]interface{} "Erreur serveur"
// @Router       /admin/profil/assignProfil [post]
func (h *livraisonHandler) AssignProfil(c *gin.Context) {
	var payload entities.PayloadAssignProfil

	// Bind JSON
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Payload invalide",
			"error":   err.Error(),
		})
		return
	}

	// Début de transaction
	tx := h.db.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Impossible de démarrer la transaction",
			"error":   tx.Error.Error(),
		})
		return
	}

	// Création en base
	if err := tx.Create(&payload).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Échec lors de l’assignation du profil à l’utilisateur",
			"error":   err.Error(),
		})
		return
	}

	// Commit
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Échec lors de la validation de la transaction",
			"error":   err.Error(),
		})
		return
	}

	// Succès
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Profil assigné à l’utilisateur avec succès",
		"data":    payload,
	})
}
