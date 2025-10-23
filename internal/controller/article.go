package controller

import (
	"car_project/internal/entities"
	"fmt"
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
// @Summary Ajouter un nouvel article
// @Description Crée un article avec ses informations et une image associée
// @Tags article
// @Accept multipart/form-data
// @Produce json
// @Param nom formData string true "Nom de l'article"
// @Param description formData string false "Description de l'article"
// @Param prix formData number true "Prix de l'article"
// @Param stock formData int true "Quantité en stock"
// @Param commercant_id formData int true "ID du commerçant"
// @Param categorie_id formData int true "ID de la catégorie"
// @Param largeur formData int false "Largeur de l'image"
// @Param hauteur formData int false "Hauteur de l'image"
// @Param ordre formData int false "Ordre de l'image"
// @Param type formData string false "Type de l'image (jpg, png, etc.)"
// @Param taille formData string false "Taille de l'image"
// @Param image formData file true "Fichier image à uploader"
// @Success 200 {object} map[string]interface{} "Article créé avec succès"
// @Failure 400 {object} map[string]interface{} "Requête invalide ou image manquante"
// @Failure 500 {object} map[string]interface{} "Erreur serveur lors de l'ajout de l'article"
// @Router /dash/article/add [post]
func (h *livraisonHandler) AjoutArticle(c *gin.Context) {
	// Récupération des champs (form-data)
	var article entities.Articles
	article.Nom = c.PostForm("nom")
	article.Description = c.PostForm("description")

	// Convertir prix et stock
	prix, _ := strconv.ParseFloat(c.PostForm("prix"), 64)
	stock, _ := strconv.Atoi(c.PostForm("stock"))
	commercantID, _ := strconv.Atoi(c.PostForm("commercant_id"))
	categorieID, _ := strconv.Atoi(c.PostForm("categorie_id"))
	largeur, _ := strconv.Atoi(c.PostForm("largeur"))
	hauteur, _ := strconv.Atoi(c.PostForm("hauteur"))
	ordre, _ := strconv.Atoi(c.PostForm("ordre"))
	imageType := c.PostForm("type")
	taille := c.PostForm("taille")

	article.Prix = prix
	article.Stock = stock
	article.Commercant_id = commercantID
	article.Categorie_id = categorieID

	// Fichier image
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Image is required",
			"error":   err.Error(),
		})
		return
	}

	// Début transaction
	tx := h.db.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to start transaction", "error": tx.Error.Error()})
		return
	}

	// Insérer l'article
	if err := tx.Create(&article).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create article", "error": err.Error()})
		return
	}

	// Définir le chemin du fichier (uploads/ + nom original)
	filePath := fmt.Sprintf("uploads/%d_%s", article.Article_id, file.Filename)

	// Sauvegarde dans uploads/
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save image", "error": err.Error()})
		return
	}

	// Insérer dans Article_Image
	imageRecord := entities.ArticleImage{
		Article_id: article.Article_id,
		Url:        filePath,
		Largeur:    largeur,
		Hauteur:    hauteur,
		Ordre:      ordre,
		Type:       imageType,
		Taille:     taille,
	}
	if err := tx.Create(&imageRecord).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save image record", "error": err.Error()})
		return
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to commit transaction", "error": err.Error()})
		return
	}

	// Succès
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Article created successfully",
		"data": gin.H{
			"article": article,
			"image":   imageRecord,
		},
	})
}

// DeleteArticle godoc
// @Summary Supprimer un article
// @Description Supprime un article en fonction de son ID
// @Tags article
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
