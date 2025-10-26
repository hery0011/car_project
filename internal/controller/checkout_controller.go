package controller

import (
	"car_project/internal/entities"
	"car_project/internal/service"
	"car_project/internal/websocket"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *livraisonHandler) Checkout(c *gin.Context) {
	// Récupérer la session depuis le contexte
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

	userID := sessionData.User.Id // ID de l'utilisateur connecté

	var payload struct {
		Address entities.Address        `json:"livraison"`
		Items   []entities.OrderItem    `json:"items"`
		Method  entities.PaymentPayload `json:"payment"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderService := service.NewOrderService(h.db)
	order, err := orderService.CreateOrder(userID, &payload.Address, payload.Items, payload.Method.Method)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Notifications WebSocket non bloquantes
	go websocket.NotifyClient(userID, order)
	// for merchantID := range order.MerchantIDs {
	// 	go websocket.NotifyMerchant(merchantID, order)
	// }
	go websocket.NotifyAdmin(order)

	c.JSON(http.StatusOK, gin.H{"order": order})
}
