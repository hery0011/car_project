package controller

import (
	"car_project/internal/config"
	"car_project/internal/entities"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// AjoutPanier crée un nouveau panier pour un client et ajoute un article dans Panier_Article.
// La création se fait dans une transaction pour assurer la cohérence.
//
//	@Summary		Créer un panier et ajouter un article
//	@Description	Crée un panier avec un status par défaut ("panier ouvert") et ajoute un article avec la quantité spécifiée.
//	@Tags			panier
//	@Accept			json
//	@Produce		json
//	@Param			panierData	body	entities.PayloadPanier	true	"Payload pour créer un panier"
//	@Success		200	{object}	map[string]interface{}	"Panier créé avec succès"
//	@Failure		400	{object}	map[string]interface{}	"Erreur de validation du payload"
//	@Failure		500	{object}	map[string]interface{}	"Erreur serveur ou transaction échouée"
//	@Router			/dash/article/panier/add [post]
func (h *livraisonHandler) AjoutPanier(c *gin.Context) {
	var payload entities.PayloadPanier

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Démarrer la transaction
	tx := h.db.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Impossible de démarrer la transaction",
		})
		return
	}

	// 1️⃣ Créer le panier
	panier := entities.Panier{
		ClientId:     payload.ClientId,
		Status_id:    config.PANIER_OUVERT,
		DateCreation: time.Now().Format("2006-01-02 15:04:05"),
	}

	if err := tx.Create(&panier).Error; err != nil {
		tx.Rollback() // Annuler la transaction
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Impossible de créer le panier",
			"error":   err.Error(),
		})
		return
	}

	// 2️⃣ Ajouter l’article dans Panier_Article
	panierArticle := entities.PanierArticle{
		PanierId:  panier.PanierId,
		ArticleId: payload.ArticleId,
		Quantite:  payload.Quantite,
	}

	if err := tx.Create(&panierArticle).Error; err != nil {
		tx.Rollback() // Annuler la transaction
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Impossible d’ajouter l’article au panier",
			"error":   err.Error(),
		})
		return
	}

	// Valider la transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Impossible de valider la transaction",
			"error":   err.Error(),
		})
		return
	}

	// ✅ Retour succès
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Panier créé avec succès",
		"panier":  panier,
		"article": panierArticle,
	})
}

// DetailPanier retourne le ou les paniers ouverts d'un client avec leurs articles.
// Un panier ouvert est défini par status_id = 4.
//
//	@Summary		Récupérer le détail du panier d'un client
//	@Description	Récupère tous les paniers ouverts (status_id = 4) pour un client donné, avec la liste des articles associés.
//	@Tags			panier
//	@Accept			json
//	@Produce		json
//	@Param			id_client	path		int	true	"ID du client"
//	@Success		200	{array}		map[string]interface{}	"Liste des paniers et articles du client"
//	@Failure		400	{object}	map[string]string	"ID client invalide"
//	@Failure		500	{object}	map[string]string	"Impossible de récupérer les paniers"
//	@Router			/dash/article/panier/{id_client}/detail [get]
func (h *livraisonHandler) DetailPanier(c *gin.Context) {
	clientIDStr := c.Param("id_client")
	clientID, err := strconv.Atoi(clientIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID client invalide"})
		return
	}

	var paniers []entities.Panier
	// Status "panier ouvert" → id_status = 4
	if err := h.db.
		Where("client_id = ? AND status_id = ?", clientID, config.PANIER_OUVERT).
		Find(&paniers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer les paniers"})
		return
	}

	// Pour chaque panier, récupérer ses articles
	var result []map[string]interface{}
	for _, panier := range paniers {
		var articles []entities.PanierArticle
		h.db.Where("panier_id = ?", panier.PanierId).Find(&articles)

		result = append(result, map[string]interface{}{
			"panier":   panier,
			"articles": articles,
		})
	}

	c.JSON(http.StatusOK, result)
}

// DeletePanier supprime un panier et ses articles associés.
// Si le panier contient encore des articles, ils seront supprimés avant (ON DELETE CASCADE ou suppression manuelle).
//
//	@Summary		Supprimer un panier
//	@Description	Supprime un panier par son ID. Si des articles sont liés, ils sont également supprimés.
//	@Tags			panier
//	@Accept			json
//	@Produce		json
//	@Param			id_panier	path		int	true	"ID du panier à supprimer"
//	@Success		200	{object}	map[string]interface{}	"Panier et articles supprimés avec succès"
//	@Failure		400	{object}	map[string]string	"ID invalide"
//	@Failure		404	{object}	map[string]string	"Panier introuvable"
//	@Failure		500	{object}	map[string]string	"Erreur lors de la suppression du panier"
//	@Router			/dash/article/panier/{id_panier}delete [delete]
func (h *livraisonHandler) DeletePanier(c *gin.Context) {
	var panier entities.Panier
	idPanier := c.Param("id_panier")

	id, err := strconv.Atoi(idPanier)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "ID invalide"})
		return
	}

	// Vérifier si le panier existe
	if err := h.db.First(&panier, "panier_id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Panier introuvable"})
		return
	}

	// Supprimer d'abord les articles liés
	if err := h.db.Where("panier_id = ?", id).Delete(&entities.PanierArticle{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Erreur lors de la suppression des articles du panier",
			"error":   err.Error(),
		})
		return
	}

	// Ensuite supprimer le panier
	if err := h.db.Delete(&panier).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Erreur lors de la suppression du panier",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Panier et ses articles supprimés avec succès",
		"data":    panier,
	})
}
