package controller

import (
	"car_project/internal/entities"
	"car_project/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *livraisonHandler) ListOrders(c *gin.Context) {
	// Récupérer la session utilisateur
	sessionInterface, exists := c.Get("sessionData")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Utilisateur non authentifié"})
		return
	}

	sessionData, ok := sessionInterface.(entities.SessionData)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer la session utilisateur"})
		return
	}

	userID := sessionData.User.Id

	// Filtrage par statut optionnel via query param
	statusCode := c.Query("status") // ex: "pending_payment", "shipped"

	orderService := service.NewOrderService(h.db)
	orders, err := orderService.ListOrders(userID, statusCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}
