package controller

import (
	"car_project/internal/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListCategorie godoc
// @Summary Liste des catégories
// @Description Récupère la liste des catégories disponibles
// @Tags categorie
// @Produce json
// @Success 200 {object} map[string]interface{} "Liste des catégories ou tableau vide"
// @Failure 400 {object} map[string]interface{} "Erreur lors de la récupération"
// @Router /dash/article/categorie/list [get]
func (h *livraisonHandler) ListCategorie(c *gin.Context) {
	var categories []entities.Categorie

	if err := h.db.Preload("SubCategories").Find(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if len(categories) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Aucune catégorie trouvée",
			"data":    []entities.Categorie{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Liste des catégories",
		"data":    categories,
	})
}
