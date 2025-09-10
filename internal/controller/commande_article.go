package controller

import (
	"car_project/internal/config"
	"car_project/internal/entities"
	"car_project/internal/ws"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ArticlePayload struct {
	ClientId       int     `json:"client_id"`
	ArticleId      int     `json:"article_id"`
	Quantite       int     `json:"quantite"`
	PrixUnitaire   float64 `json:"prix_unitaire"`
	LieuxLivraison string  `json:"lieuxLivraison" gorm:"column:lieuxLivraison"`
	Latitude       float64 `json:"latitude" gorm:"latitude"`
	Longitude      float64 `json:"longitude" gorm:"longitude"`
}

// AjoutCommande godoc
// @Summary Ajouter une commande
// @Description Crée une nouvelle commande avec ses articles et envoie une notification WS aux commerçants concernés
// @Tags commande
// @Accept json
// @Produce json
// @Param payload body []controller.ArticlePayload true "Liste des articles de la commande"
// @Success 200 {object} map[string]interface{} "Commande créée avec succès, contient la commande et les articles"
// @Failure 400 {object} map[string]interface{} "Payload vide ou erreur lors de l'insertion"
// @Failure 500 {object} map[string]interface{} "Erreur serveur lors de la création de la commande"
// @Router /dash/article/commande/add [post]
func (h *livraisonHandler) AjoutCommande(c *gin.Context) {
	var payload ArticlePayload

	// Lire le JSON venant du frontend
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	// Vérifier si le payload est vide (ex: quantite <= 0 ou client_id manquant)
	if payload.ClientId == 0 || payload.ArticleId == 0 || payload.Quantite == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "payload invalide ou incomplet",
		})
		return
	}

	// Calcul du montant total
	montantTotal := float64(payload.Quantite) * payload.PrixUnitaire

	// Créer la commande
	commande := entities.Commande{
		ClientId:       payload.ClientId,
		DateCommande:   time.Now().Format("2006-01-02"),
		MontantTotal:   montantTotal,
		StatusId:       config.COMMANDE_OUVERT,
		LivreurAssign:  config.STATUS_ASSIGN,
		LieuxLivraison: payload.LieuxLivraison,
		Latitude:       payload.Latitude,
		Longitude:      payload.Longitude,
	}

	if err := h.db.Create(&commande).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	// Insérer l'article lié
	articleCommande := entities.CommandeArticle{
		CommandeId:   commande.CommandeId,
		ArticleId:    payload.ArticleId,
		Quantite:     payload.Quantite,
		PrixUnitaire: payload.PrixUnitaire,
	}

	if err := h.db.Create(&articleCommande).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	// 🔹 Envoyer notification WS au commerçant concerné
	var article entities.Articles
	if err := h.db.First(&article, payload.ArticleId).Error; err == nil {
		msg := fmt.Sprintf("Nouvelle commande %d : Article %d x%d",
			commande.CommandeId, payload.ArticleId, payload.Quantite)
		fmt.Println("Envoi notif à commerçant:", article.Commercant_id, "msg:", msg)
		ws.NotifyCommercant(article.Commercant_id, msg)
	}

	// Réponse finale
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"message":  "Commande créée avec succès",
		"commande": commande,
		"article":  articleCommande,
	})
}

// AssignCommande godoc
// @Summary Assigner une commande à un livreur
// @Description Met à jour une commande en assignant un livreur et en changeant son statut en "EN COURS"
// @Tags commande
// @Accept json
// @Produce json
// @Param id_commande path int true "ID de la commande"
// @Param id_livreur path int true "ID du livreur"
// @Success 200 {object} map[string]interface{} "Commande assignée avec succès"
// @Failure 400 {object} map[string]interface{} "ID commande ou livreur invalide"
// @Failure 500 {object} map[string]interface{} "Erreur serveur lors de l'assignation"
// @Router /dash/article/commande/{id_commande}/assign/{id_livreur} [put]
func (h *livraisonHandler) AssignCommande(c *gin.Context) {
	var entiteCommande entities.Commande
	idCommande := c.Param("id_commande")
	idLivreur := c.Param("id_livreur")

	idC, err := strconv.Atoi(idCommande)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ID commande invalide",
		})
		return
	}

	idL, err := strconv.Atoi(idLivreur)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ID livreur invalide",
		})
		return
	}

	// Mise à jour du champ livreur_assign et status
	if err := h.db.Model(&entiteCommande).
		Where("commande_id = ?", idC).
		Updates(map[string]interface{}{
			"livreur_assign": idL,
			"status_id":      config.COMMANDE_EN_COURS,
		}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Erreur lors de l'assignation",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Commande assignée avec succès",
	})
}

// ListeCommandeOuvert godoc
// @Summary Liste les commandes ouvertes
// @Description Récupère toutes les commandes dont le statut est "ouvertes", avec les informations du client et des articles associés
// @Tags commande
// @Produce json
// @Success 200 {object} map[string]interface{} "Liste des commandes ouvertes"
// @Failure 400 {object} map[string]interface{} "Erreur lors de la récupération des commandes"
// @Router /dash/article/commande/commandeOuvert [get]
func (h *livraisonHandler) ListeCommandeOuvert(c *gin.Context) {
	type Result struct {
		CommandeId        int     `json:"commande_id"`
		DateCommande      string  `json:"date_commande"`
		MontantTotal      float64 `json:"montant_total"`
		ClientNom         string  `json:"client_nom"`
		ClientPrenom      string  `json:"client_prenom"`
		ClientEmail       string  `json:"client_email"`
		Adresse           string  `json:"adresse"`
		CommandeArticleId int     `json:"commande_article_id"`
		ArticleId         int     `json:"article_id"`
		Quantite          int     `json:"quantite"`
		PrixUnitaire      float64 `json:"prix_unitaire"`
		LieuxLivraison    string  `json:"lieux_livraison"`
		Latitude          float64 `json:"latitude"`
		Longitude         float64 `json:"longitude"`
	}

	var results []Result

	err := h.db.Table("Commande c").
		Select(`c.commande_id, c.date_commande, c.montant_total, c.lieux_livraison, c.latitude, c.longitude,
				cl.nom AS client_nom, cl.prenom AS client_prenom, cl.email AS client_email, cl.adresse AS adresse,
				ca.commande_article_id, ca.article_id, ca.quantite, ca.prix_unitaire`).
		Joins("JOIN Client cl ON c.client_id = cl.client_id").
		Joins("JOIN Commande_Article ca ON c.commande_id = ca.commande_id").
		Where("c.status_id = ?", config.COMMANDE_OUVERT).
		Scan(&results).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "liste de commandes ouvertes",
		"data":    results,
	})
}

// ListeCommandeAssign godoc
// @Summary Liste les commandes assignées à un livreur
// @Description Récupère toutes les commandes ouvertes assignées à un livreur spécifique, avec les informations du client et des articles associés
// @Tags commande
// @Produce json
// @Param user_id path int true "ID du livreur"
// @Success 200 {object} map[string]interface{} "Liste des commandes assignées"
// @Failure 400 {object} map[string]interface{} "ID utilisateur invalide ou erreur lors de la récupération"
// @Failure 500 {object} map[string]interface{} "Erreur serveur"
// @Router /dash/article/commande/commandeAssign/{user_id} [get]
func (h *livraisonHandler) ListeCommandeAssign(c *gin.Context) {
	// Récupération et conversion de l'ID utilisateur
	userIDParam := c.Param("user_id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ID utilisateur invalide",
		})
		return
	}

	// Struct pour le résultat
	type Result struct {
		CommandeId        int     `json:"commande_id"`
		DateCommande      string  `json:"date_commande"`
		MontantTotal      float64 `json:"montant_total"`
		ClientNom         string  `json:"client_nom"`
		ClientPrenom      string  `json:"client_prenom"`
		ClientEmail       string  `json:"client_email"`
		Adresse           string  `json:"adresse"`
		CommandeArticleId int     `json:"commande_article_id"`
		ArticleId         int     `json:"article_id"`
		Quantite          int     `json:"quantite"`
		PrixUnitaire      float64 `json:"prix_unitaire"`
	}

	var results []Result

	// Requête GORM
	err = h.db.Table("Commande c").
		Select(`c.commande_id, c.date_commande, c.montant_total,
				cl.nom AS client_nom, cl.prenom AS client_prenom, cl.email AS client_email, cl.adresse AS adresse,
				ca.commande_article_id, ca.article_id, ca.quantite, ca.prix_unitaire`).
		Joins("JOIN Client cl ON c.client_id = cl.client_id").
		Joins("JOIN Commande_Article ca ON c.commande_id = ca.commande_id").
		Where("c.status_id = ? AND c.livreur_assign = ?", config.COMMANDE_EN_COURS, userID).
		Scan(&results).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Erreur lors de la récupération des commandes : " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Liste des commandes assignées à l'utilisateur",
		"data":    results,
	})
}
