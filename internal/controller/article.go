package controller

import (
	"car_project/internal/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetArticleImages godoc
// @Summary Récupérer les images d'articles
// @Description Retourne une liste paginée d'images (5 par page).
// @Tags article
// @Accept json
// @Produce json
// @Param page query int true "Numéro de la page (commence à 1)"
// @Success 200 {object} entities.ArticleResponse
// @Failure 400 {object} map[string]string
// @Router /dash/article/list [get]
func (h *livraisonHandler) ListArticle(c *gin.Context) {
	// Récupérer le numéro de page depuis la query string (par défaut = 1)
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit := 5
	offset := (page - 1) * limit

	var articles []entities.Article
	var total int64

	// Compter le nombre total d'articles
	if err := h.db.Model(&entities.Article{}).Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Erreur lors du comptage des articles",
			"error":   err.Error(),
		})
		return
	}

	// Charger les données avec pagination + Preload
	if err := h.db.
		Preload("Images").
		Preload("Categorie").
		Preload("Commercant").
		Limit(limit).
		Offset(offset).
		Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Erreur lors de la récupération des articles",
			"error":   err.Error(),
		})
		return
	}

	// Vérifier si aucun article trouvé
	if len(articles) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Aucun article trouvé pour cette page",
			"data":    []entities.ArticleResponse{},
		})
		return
	}

	// Transformer en ArticleResponse
	var response []entities.ArticleResponse
	for _, a := range articles {
		resp := entities.ArticleResponse{
			ArticleID:   a.Article_id,
			Nom:         a.Nom,
			Description: a.Description,
			Prix:        a.Prix,
			Stock:       a.Stock,
			Categorie:   a.Categorie,
			Commercant:  a.Commercant,
			Images:      a.Images,
		}
		response = append(response, resp)
	}

	// Calcul du nombre total de pages
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	// Réponse finale
	c.JSON(http.StatusOK, gin.H{
		"status":     http.StatusOK,
		"message":    "Liste des articles récupérée avec succès",
		"page":       page,
		"limit":      limit,
		"totalItems": total,
		"totalPages": totalPages,
		"count":      len(response),
		"data":       response,
	})
}

// AjoutArticle godoc
// @Summary Ajouter un article
// @Description Crée un nouvel article dans la base de données
// @Tags article
// @Accept  json
// @Produce  json
// @Param   article body entities.Articles true "Détails de l'article"
// @Success 200 {object} map[string]interface{} "Article créé avec succès"
// @Failure 400 {object} map[string]interface{} "Payload invalide"
// @Failure 500 {object} map[string]interface{} "Erreur serveur"
// @Router /dash/article/add [post]
func (h *livraisonHandler) AjoutArticle(c *gin.Context) {
	var payload entities.Articles

	// Vérifier que le JSON est correct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
		return
	}

	// Vérifier la connexion DB
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Database connection not initialized",
		})
		return
	}

	// Début transaction
	tx := h.db.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to start transaction",
			"error":   tx.Error.Error(),
		})
		return
	}

	// Création de l'article
	if err := tx.Create(&payload).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to create article",
			"error":   err.Error(),
		})
		return
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to commit transaction",
			"error":   err.Error(),
		})
		return
	}

	// Succès
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Article created successfully",
		"data":    payload,
	})
}

// DeleteArticle godoc
// @Summary Supprimer un article
// @Description Supprime un article en fonction de son ID
// @Tags Articles
// @Param id path int true "ID de l'article"
// @Produce  json
// @Success 200 {object} map[string]interface{} "Article supprimé avec succès"
// @Failure 400 {object} map[string]interface{} "ID invalide"
// @Failure 404 {object} map[string]interface{} "Article introuvable"
// @Failure 500 {object} map[string]interface{} "Erreur serveur"
// @Router /dash/article/{id}/delete [delete]
func (h *livraisonHandler) DeleteArticle(c *gin.Context) {
	var article entities.Articles
	id := c.Param("id")

	// Conversion de l'ID
	idArticle, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid article ID",
			"error":   err.Error(),
		})
		return
	}

	// Suppression
	result := h.db.Where("article_id = ?", idArticle).Delete(&article)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to delete article",
			"error":   result.Error.Error(),
		})
		return
	}

	// Vérifier si un article a été supprimé
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Article not found",
		})
		return
	}

	// Succès
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Article deleted successfully",
		"id":      idArticle,
	})
}
