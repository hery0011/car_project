package controller

import (
	"car_project/internal/entities"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"car_project/internal/helper"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	limit := 13
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
		// Preload("Categorie").
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
			ArticleID:   a.ArticleID,
			Nom:         a.Nom,
			Slug:        a.Slug,
			Description: a.Description,
			Prix:        a.Prix,
			Stock:       a.Stock,
			// Categorie:   a.Categories[],
			Commercant: a.Commercant,
			Images:     a.Images,
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

// GetArticleImages godoc
// @Summary Récupérer les images d'articles par commerçant
// @Description Retourne une liste paginée d'articles appartenant à un commerçant spécifique.
// @Tags article
// @Accept json
// @Produce json
// @Param idCommercant path int true "ID du commerçant"
// @Param page query int false "Numéro de la page (commence à 1)"
// @Success 200 {object} entities.ArticleResponse
// @Failure 400 {object} map[string]string
// @Router /dash/article/list/{idCommercant} [get]
func (h *livraisonHandler) ListeArticleByCommercant(c *gin.Context) {
	// 🔹 Récupérer l'ID commerçant depuis l'URL
	idCommercantStr := c.Param("idCommercant")
	idCommercant, err := strconv.Atoi(idCommercantStr)
	if err != nil || idCommercant <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ID commerçant invalide",
		})
		return
	}

	// 🔹 Récupérer la page
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit := 10
	offset := (page - 1) * limit

	var articles []entities.Article
	var total int64

	// 🔹 Compter le nombre total d'articles du commerçant
	if err := h.db.Model(&entities.Article{}).
		Where("commercant_id = ?", idCommercant).
		Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Erreur lors du comptage des articles",
			"error":   err.Error(),
		})
		return
	}

	// 🔹 Charger les articles filtrés par commerçant avec pagination
	if err := h.db.
		Preload("Images").
		Preload("Categorie").
		Preload("Commercant").
		Where("commercant_id = ?", idCommercant).
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

	// 🔹 Vérifier si aucun article trouvé
	if len(articles) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Aucun article trouvé pour ce commerçant",
			"data":    []entities.ArticleResponse{},
		})
		return
	}

	// 🔹 Transformer en ArticleResponse
	var response []entities.ArticleResponse
	for _, a := range articles {
		response = append(response, entities.ArticleResponse{
			ArticleID:   a.ArticleID,
			Nom:         a.Nom,
			Description: a.Description,
			Prix:        a.Prix,
			Stock:       a.Stock,
			Categorie:   a.Categories,
			Commercant:  a.Commercant,
			Images:      a.Images,
		})
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	// 🔹 Réponse finale
	c.JSON(http.StatusOK, gin.H{
		"status":     http.StatusOK,
		"message":    "Articles du commerçant récupérés avec succès",
		"page":       page,
		"limit":      limit,
		"totalItems": total,
		"totalPages": totalPages,
		"count":      len(response),
		"data":       response,
	})
}

// GetArticleDetail godoc
// @Summary Récupérer les détails d'un article
// @Description Retourne les informations détaillées d'un article (images, catégorie, commerçant)
// @Tags article
// @Accept json
// @Produce json
// @Param id path int true "ID de l'article"
// @Success 200 {object} entities.ArticleResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /dash/article/{id} [get]
func (h *livraisonHandler) GetArticleDetail(c *gin.Context) {
	// Récupérer l'ID depuis l'URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ID d'article invalide",
		})
		return
	}

	var article entities.Article

	// Rechercher l'article avec les relations
	if err := h.db.
		Preload("Images").
		Preload("Categories").
		Preload("Commercant").
		First(&article, "article_id = ?", id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Article non trouvé",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Erreur lors de la récupération de l'article",
				"error":   err.Error(),
			})
		}
		return
	}

	// Transformer en ArticleResponse
	response := entities.ArticleResponse{
		ArticleID:   article.ArticleID,
		Nom:         article.Nom,
		Description: article.Description,
		Prix:        article.Prix,
		Stock:       article.Stock,
		Categorie:   article.Categories,
		Commercant:  article.Commercant,
		Images:      article.Images,
	}

	// Réponse finale
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Détails de l'article récupérés avec succès",
		"data":    response,
	})
}

func (h *livraisonHandler) ListCategories(c *gin.Context) {
	var categories []entities.Categorie

	// Charger toutes les catégories
	if err := h.db.Preload("SubCategories").Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Erreur lors de la récupération des catégories",
			"error":   err.Error(),
		})
		return
	}

	// Construire la map pour accès rapide par ID
	categoryMap := make(map[int]*entities.CategoryResponse)
	var roots []*entities.CategoryResponse

	for _, cat := range categories {
		categoryMap[cat.Categorie_id] = &entities.CategoryResponse{
			CategoryId:    uint(cat.Categorie_id),
			Nom:           cat.Nom,
			ImageUrl:      cat.ImageUrl,
			SubCategories: []*entities.CategoryResponse{},
		}
	}

	// Construire la hiérarchie parent/enfant
	for _, cat := range categories {
		cr := categoryMap[cat.Categorie_id]
		if cat.Parent_id != 0 {
			parent := categoryMap[cat.Parent_id]
			if parent != nil {
				parent.SubCategories = append(parent.SubCategories, cr)
			}
		} else {
			roots = append(roots, cr)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Liste des catégories récupérée avec succès",
		"count":   len(roots),
		"data":    roots,
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
// @Param image formData string true "Image encodée en base64"
// @Success 200 {object} map[string]interface{} "Article créé avec succès"
// @Failure 400 {object} map[string]interface{} "Requête invalide ou image manquante"
// @Failure 500 {object} map[string]interface{} "Erreur serveur lors de l'ajout de l'article"
// @Router /dash/article/add [post]
func (h *livraisonHandler) AjoutArticle(c *gin.Context) {
	var input entities.ArticleCreateRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Erreur de Binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Erreur de validation de la requête JSON",
			"error":   err.Error(),
		})
		return
	}

	if input.CommercantID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "CommercantID est manquant ou invalide (doit être > 0)."})
		return
	}
	var existingCommercant entities.Commercant
	if h.db.First(&existingCommercant, input.CommercantID).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": fmt.Sprintf("Commercant ID %d non trouvé. Impossible d'associer l'article.", input.CommercantID),
		})
		return
	}

	article := entities.Article{
		Nom:              input.Nom,
		Slug:             generateSlug(input.Nom),
		ShortDescription: "",
		Description:      input.Description,
		Status:           "draft",
		IsActive:         true,
		Prix:             input.Prix,
		Stock:            input.Stock,
		CommercantID:     input.CommercantID,
	}

	tx := h.db.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to start transaction", "error": tx.Error.Error()})
		return
	}

	if err := tx.Create(&article).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create article (Clé étrangère CommercantID)", "error": err.Error()})
		return
	}

	articleID := article.ArticleID

	var savedImages []entities.ArticleImage
	for i, imgPayload := range input.Images {
		base64Image := imgPayload.Base64Data
		if coI := strings.Index(base64Image, ","); coI != -1 {
			base64Image = base64Image[coI+1:]
		}
		imgData, err := base64.StdEncoding.DecodeString(base64Image)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Image #%d: Invalid base64 data", i+1), "error": err.Error()})
			return
		}

		// 🔑 Gestion des chemins et création récursive des dossiers
		fileName := fmt.Sprintf("/uploads/commercants/%d/articles/%d/uploads/%d-%d-%d.jpg", input.CommercantID, articleID, articleID, i, time.Now().UnixNano())

		dirPath := filepath.Dir(fileName)

		if err := os.MkdirAll("."+dirPath, os.ModePerm); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Image #%d: Failed to create directory structure %s", i+1, dirPath), "error": err.Error()})
			return
		}

		if err := ioutil.WriteFile("."+fileName, imgData, 0644); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Image #%d: Failed to save file", i+1), "error": err.Error()})
			return
		}

		imageRecord := entities.ArticleImage{
			Article_id: articleID,
			Url:        fileName,
			Largeur:    imgPayload.Largeur,
			Hauteur:    imgPayload.Hauteur,
			Ordre:      imgPayload.Ordre,
			Type:       imgPayload.Type,
			Taille:     imgPayload.Taille,
		}
		if err := tx.Create(&imageRecord).Error; err != nil {
			os.Remove("." + fileName)
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Image #%d: Failed to save record", i+1), "error": err.Error()})
			return
		}
		savedImages = append(savedImages, imageRecord)
	}

	var linkedCategories []entities.Categorie
	for _, catID := range input.CategorieIDs {
		var count int64
		if err := tx.Model(&entities.Categorie{}).Where("categorie_id = ?", catID).Count(&count).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error during category check", "error": err.Error()})
			return
		}

		if count == 0 {
			tx.Rollback()
			c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Category ID %d not found", catID)})
			return
		}

		if err := tx.Exec("INSERT INTO article_category (article_id, categorie_id) VALUES (?, ?)", articleID, catID).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to manually link category", "error": err.Error()})
			return
		}

		var cat entities.Categorie
		h.db.First(&cat, catID)
		linkedCategories = append(linkedCategories, cat)
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to commit transaction", "error": err.Error()})
		return
	}

	article.Images = savedImages
	article.Categories = linkedCategories

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Article created successfully",
		"data":    article,
	})
}

func generateSlug(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "-")
	return s
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

func (h *livraisonHandler) FilterArticleByCommercant(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit := 10
	offset := (page - 1) * limit

	// Récupérer le nom du commerçant depuis l'URL
	commercantNom := c.Param("commercant")

	var articles []entities.Article
	var total int64

	// Base de la requête
	query := h.db.Model(&entities.Article{}).
		Preload("Images").
		Preload("Categorie").
		Preload("Commercant")

	// Si un nom de commerçant est fourni
	if commercantNom != "" {
		likeValue := fmt.Sprintf("%%%s%%", commercantNom)
		query = query.
			Joins("JOIN commercant ON commercant.commercant_id = article.commercant_id").
			Where("commercant.nom LIKE ?", likeValue)
	}

	// 🔑 CORRECTION CLÉ : Utiliser Session() pour le comptage.
	// Cela crée une copie de la requête avec les JOINs et WHERE pour obtenir le total FILTRÉ,
	// sans modifier l'état de la requête 'query' originale.
	countQuery := query.Session(&gorm.Session{})

	// Compter le total filtré
	if err := countQuery.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Erreur lors du comptage des articles",
			"error":   err.Error(),
		})
		return
	}

	// Charger les données paginées
	// La requête 'query' initiale conserve son filtre (JOIN/WHERE) intact
	if err := query.
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

	if len(articles) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Aucun article trouvé pour cette page",
			"data":    []entities.ArticleResponse{},
		})
		return
	}

	// Mapper les résultats
	var response []entities.ArticleResponse
	for _, a := range articles {
		response = append(response, entities.ArticleResponse{
			ArticleID:   a.ArticleID,
			Nom:         a.Nom,
			Description: a.Description,
			Prix:        a.Prix,
			Stock:       a.Stock,
			Categorie:   a.Categories,
			Commercant:  a.Commercant,
			Images:      a.Images,
		})
	}

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

func (h *livraisonHandler) FilterArticleByName(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit := 10
	offset := (page - 1) * limit

	// Récupérer le nom de l'article depuis la QUERY STRING (ex: ?article=pain)
	articleNom := c.Param("article")
	fmt.Println(articleNom)

	var articles []entities.Article
	var total int64

	// Base de la requête
	query := h.db.Model(&entities.Article{}).
		Preload("Images").
		Preload("Categorie").
		Preload("Commercant")

	// Si un nom d'article est fourni, on filtre.
	// ⚠️ PAS BESOIN DE JOIN, car le champ 'Nom' est directement dans la table 'article'.
	if articleNom != "" {
		likeValue := fmt.Sprintf("%%%s%%", articleNom)
		query = query.
			Where("article.nom LIKE ?", likeValue) // 👈 CHANGEMENT : Filtre sur article.nom
	}

	// Filtre par Nom d'Article (champ direct)
	if articleNom != "" {
		likeValue := fmt.Sprintf("%%%s%%", articleNom)
		// Note: Si 'commercantNom' est aussi présent, GORM ajoute cette clause WHERE avec le JOIN existant.
		query = query.Where("article.nom LIKE ?", likeValue)
	}
	// 🔑 CORRECTION CLÉ : Utiliser Session() pour le comptage (maintient l'état du filtre).
	countQuery := query.Session(&gorm.Session{})

	// Compter le total filtré
	if err := countQuery.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Erreur lors du comptage des articles",
			"error":   err.Error(),
		})
		return
	}

	// Charger les données paginées
	if err := query.
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

	if len(articles) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Aucun article trouvé pour cette page",
			"data":    []entities.ArticleResponse{},
		})
		return
	}

	// Mapper les résultats (inchangé)
	var response []entities.ArticleResponse
	for _, a := range articles {
		response = append(response, entities.ArticleResponse{
			ArticleID:   a.ArticleID,
			Nom:         a.Nom,
			Description: a.Description,
			Prix:        a.Prix,
			Stock:       a.Stock,
			Categorie:   a.Categories,
			Commercant:  a.Commercant,
			Images:      a.Images,
		})
	}

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
func (h *livraisonHandler) FilterArticleByCategorie(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit := 10
	offset := (page - 1) * limit

	// Récupérer le nom de la catégorie depuis la QUERY STRING (ex: ?categorie=boissons)
	categorieNom := c.Param("categorie")

	var articles []entities.Article
	var total int64

	// Base de la requête
	query := h.db.Model(&entities.Article{}).
		Preload("Images").
		Preload("Categorie").
		Preload("Commercant")

	// Si un nom de catégorie est fourni, on filtre.
	if categorieNom != "" {
		likeValue := fmt.Sprintf("%%%s%%", categorieNom)

		// 🔑 CORRECTION CLÉ : Utiliser Joins pour filtrer sur une table associée (Categorie)
		query = query.
			Joins("JOIN categorie ON categorie.categorie_id = article.categorie_id").
			Where("categorie.nom LIKE ?", likeValue)
	}

	// 🔑 Utiliser Session() pour le comptage (isole l'état du filtre).
	countQuery := query.Session(&gorm.Session{})

	// Compter le total filtré
	if err := countQuery.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Erreur lors du comptage des articles",
			"error":   err.Error(),
		})
		return
	}

	// Charger les données paginées
	if err := query.
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

	if len(articles) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Aucun article trouvé pour cette page",
			"data":    []entities.ArticleResponse{},
		})
		return
	}

	// Mapper les résultats (inchangé)
	var response []entities.ArticleResponse
	for _, a := range articles {
		response = append(response, entities.ArticleResponse{
			ArticleID:   a.ArticleID,
			Nom:         a.Nom,
			Description: a.Description,
			Prix:        a.Prix,
			Stock:       a.Stock,
			Categorie:   a.Categories,
			Commercant:  a.Commercant,
			Images:      a.Images,
		})
	}

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

func (h *livraisonHandler) FilterArticles(c *gin.Context) {
	// --- Récupérer l'utilisateur connecté ---
	userID, err := helper.GetUserID(c)
	if err != nil {
		// l'erreur a déjà été gérée dans GetUserID, on stoppe le handler
		return
	}

	// --- Pagination ---
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	limit := 10
	offset := (page - 1) * limit

	// --- Filtre optionnel par nom et catégorie ---
	name := c.Query("name")
	// categorie := c.Query("categorie")

	var articles []entities.Article
	var total int64

	// --- Base query ---
	query := h.db.Model(&entities.Article{}).
		Preload("Images").
		// Preload("Categorie").
		Preload("Commercant")

	// --- Filtrer uniquement les articles du commerçant de l'utilisateur connecté ---
	var user entities.User
	if err := h.db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Impossible de récupérer l'utilisateur",
			"error":   err.Error(),
		})
		return
	}

	if user.CommercantID != nil {
		query = query.Where("article.commercant_id = ?", *user.CommercantID)
	}

	// --- Autres filtres ---
	if name != "" {
		query = query.Where("article.nom LIKE ?", "%"+name+"%")
	}

	// if categorie != "" {
	// 	query = query.Joins("JOIN categorie ON categorie.categorie_id = article.categorie_id").
	// 		Where("categorie.nom LIKE ?", "%"+categorie+"%")
	// }

	// --- Count total filtered ---
	countQuery := query.Session(&gorm.Session{})
	if err := countQuery.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Erreur lors du comptage des articles",
			"error":   err.Error(),
		})
		return
	}

	// --- Pagination & récupération des données ---
	if err := query.Limit(limit).Offset(offset).Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Erreur lors de la récupération des articles",
			"error":   err.Error(),
		})
		return
	}

	if len(articles) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Aucun article trouvé pour cette page",
			"data":    []entities.ArticleResponse{},
		})
		return
	}

	// --- Mapping vers ArticleResponse ---
	response := make([]entities.ArticleResponse, 0, len(articles))
	for _, a := range articles {
		response = append(response, entities.ArticleResponse{
			ArticleID:   a.ArticleID,
			Nom:         a.Nom,
			Description: a.Description,
			Prix:        a.Prix,
			Stock:       a.Stock,
			Categorie:   a.Categories,
			Commercant:  a.Commercant,
			Images:      a.Images,
		})
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	// --- Réponse finale ---
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
