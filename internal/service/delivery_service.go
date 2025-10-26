package service

import (
	"errors"
	"fmt"
	"time"

	"car_project/internal/entities"

	"gorm.io/gorm"
)

type DeliveryService struct {
	db *gorm.DB
}

func NewDeliveryService(db *gorm.DB) *DeliveryService {
	return &DeliveryService{db: db}
}

// CreateTicketFromOrder crée un ticket de livraison à partir d'une commande existante.
func (s *DeliveryService) CreateTicketFromOrder(ord *entities.Order, createdBy int) error {
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Recharger order complet si nécessaire (items, status, addresses)
	var order entities.Order
	if ord != nil && ord.OrderID != 0 {
		if err := tx.
			Preload("Items").
			Preload("Status").
			Preload("OrderAddresses"). // preload table liaison
			First(&order, ord.OrderID).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		tx.Rollback()
		return errors.New("order information required")
	}

	// Vérifier que les adresses sont présentes
	if len(order.OrderAddresses) == 0 {
		tx.Rollback()
		return fmt.Errorf("order %d has no addresses linked", order.OrderID)
	}

	// On prend le premier lien (si plusieurs, adapter selon besoin)
	addrLink := order.OrderAddresses[0]
	val := 3000.0

	now := time.Now()

	ticket := &entities.DeliveryTicket{
		OrderID:            order.OrderID,
		NomTicket:          "Commande - Article",
		ClientID:           createdBy,
		PickupAddressID:    addrLink.AdresseID,
		DropoffAddressID:   addrLink.AdresseID,
		DeliveryPrice:      &val, // nil par défaut, admin/livreur pourra mettre à jour
		PriceLastUpdatedBy: nil,  // système par défaut
		StatusID:           s.getPendingStatusID(tx),
		AssignedTo:         nil,
		CreatedAt:          now.Format("2006-10-26 12:13:15"),
		UpdatedAt:          now.Format("2006-10-26 12:13:15"),
	}

	if err := tx.Create(ticket).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// helper pour obtenir l'id du statut 'pending' dans delivery_ticket_status
func (s *DeliveryService) getPendingStatusID(tx *gorm.DB) int {
	var st entities.DeliveryTicketStatus
	if err := tx.Where("code = ?", "pending").First(&st).Error; err != nil {
		fmt.Println("-------------------...", err)
		st = entities.DeliveryTicketStatus{
			Code:    "pending",
			Label:   "En attente",
			IsFinal: false,
		}
		tx.Create(&st)
	}
	return st.ID
}
