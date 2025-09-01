package controller

import (
	"car_project/internal/config"
	"car_project/internal/entities"
	"car_project/internal/ws"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ArticlePayload struct {
	ClientId     int     `json:"client_id"`
	ArticleId    int     `json:"article_id"`
	Quantite     int     `json:"quantite"`
	PrixUnitaire float64 `json:"prix_unitaire"`
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
// @Router /ddash/article/commande/add [post]
func (h *livraisonHandler) AjoutCommande(c *gin.Context) {

	var payload []ArticlePayload

	// Lire le JSON venant du frontend
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": err.Error(),
		})
		return
	}

	if len(payload) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusOK,
			"message": "payload vide",
		})
		return
	}

	// Calcul du montant total
	var montantTotal float64
	for _, item := range payload {
		montantTotal += float64(item.Quantite) * item.PrixUnitaire
	}

	// Créer la commande
	commande := entities.Commande{
		ClientId:     payload[0].ClientId, // même client pour tous les articles
		DateCommande: time.Now().Format("2006-01-02"),
		MontantTotal: montantTotal,
		StatusId:     config.COMMANDE_OUVERT,
	}

	if err := h.db.Create(&commande).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data":    http.StatusOK,
			"message": err.Error(),
		})
		return
	}

	// Insérer les articles liés
	var articles []entities.CommandeArticle
	for _, item := range payload {
		articles = append(articles, entities.CommandeArticle{
			CommandeId:   commande.CommandeId,
			ArticleId:    item.ArticleId,
			Quantite:     item.Quantite,
			PrixUnitaire: item.PrixUnitaire,
		})
	}

	if err := h.db.Create(&articles).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusOK,
			"message": err.Error(),
		})
		return
	}

	// 🔹 Envoyer notification WS aux commerçants concernés
	for _, item := range articles {
		var article entities.Articles
		if err := h.db.First(&article, item.ArticleId).Error; err == nil {
			msg := fmt.Sprintf("Nouvelle commande %d : Article %d x%d", item.CommandeId, item.ArticleId, item.Quantite)
			fmt.Println("Envoi notif à commerçant:", article.Commercant_id, "msg:", msg)
			ws.NotifyCommercant(article.Commercant_id, msg)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"message":  "Commande créée avec succès",
		"commande": commande,
		"articles": articles,
	})
}
