package controller

import (
	"net/http"

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

type UpdateTicketInput struct {
	NomTicket     string   `json:"nom_ticket"`
	DeliveryPrice *float64 `json:"delivery_price"`
	StatusID      int      `json:"status_id"`
}

func (h *livraisonHandler) UpdateTicket(c *gin.Context) {
	var input UpdateTicketInput
	id := c.Param("id")

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var ticket entities.DeliveryTicket
	if err := h.db.First(&ticket, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ticket not found"})
		return
	}

	ticket.NomTicket = input.NomTicket
	ticket.DeliveryPrice = input.DeliveryPrice
	ticket.StatusID = input.StatusID

	if err := h.db.Save(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

type AssignTicketInput struct {
	AssignedTo int `json:"assigned_to"`
}

func (h *livraisonHandler) AssignTicket(c *gin.Context) {

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

	var input AssignTicketInput
	id := c.Param("id")

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var ticket entities.DeliveryTicket
	if err := h.db.First(&ticket, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ticket not found"})
		return
	}

	ticket.AssignedTo = &userID

	var status entities.DeliveryTicketStatus
	if err := h.db.Where("code = ?", "assigned").First(&status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer le statut assigned"})
		return
	}
	ticket.StatusID = status.ID

	if err := h.db.Save(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ticket)
}
