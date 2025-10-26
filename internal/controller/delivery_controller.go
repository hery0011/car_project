package controller

import (
	"net/http"
	"strconv"

	"car_project/internal/entities"
	"car_project/internal/service"

	"github.com/gin-gonic/gin"
)

type DeliveryHandler struct {
	service *service.DeliveryService
}

func NewDeliveryHandler(s *service.DeliveryService) *DeliveryHandler {
	return &DeliveryHandler{service: s}
}

// POST /tickets
func (h *DeliveryHandler) CreateTicket(c *gin.Context) {
	var payload entities.DeliveryTicket
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateTicket(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payload)
}

func (h *livraisonHandler) GetTickets(c *gin.Context) {
	var tickets []entities.DeliveryTicket
	// On preload les adresses et le statut
	if err := h.db.
		Preload("PickupAddress").
		Preload("DropoffAddress").
		Preload("Status").
		Find(&tickets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

// PUT /tickets/:id
func (h *DeliveryHandler) UpdateTicket(c *gin.Context) {
	ticketID, _ := strconv.Atoi(c.Param("id"))
	var payload map[string]interface{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateTicket(ticketID, payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket mis à jour"})
}
