package controller

import (
	"car_project/internal/entities"
	"car_project/internal/service"
	"car_project/internal/websocket"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *livraisonHandler) Checkout(c *gin.Context) {
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

	var payload struct {
		PickupAddress  entities.Address        `json:"pickup_address"`
		DropoffAddress entities.Address        `json:"dropoff_address"`
		Items          []entities.OrderItem    `json:"items"`
		Method         entities.PaymentPayload `json:"payment"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderService := service.NewOrderService(h.db)
	order, err := orderService.CreateOrder(
		userID,
		&payload.DropoffAddress,
		payload.Items,
		payload.Method.Method,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// --- Notifications WebSocket ---
	go websocket.NotifyClient(userID, order)
	go websocket.NotifyAdmin(order)

	// --- Création non bloquante du ticket de livraison ---
	deliveryService := service.NewDeliveryService(h.db)
	go func(ord *entities.Order, uid int) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic during CreateTicketFromOrder: %v", r)
			}
		}()

		var err error
		for attempt := 1; attempt <= 2; attempt++ {
			err = deliveryService.CreateTicketFromOrder(ord, uid)
			if err == nil {
				// TODO: notifier livreurs/admins que ticket créé
				return
			}
			log.Printf("CreateTicketFromOrder attempt %d failed: %v", attempt, err)
			time.Sleep(time.Second * time.Duration(attempt))
		}
		log.Printf("CreateTicketFromOrder failed after retries: %v", err)
	}(order, userID)

	c.JSON(http.StatusOK, gin.H{"order": order})
}
