package service

import (
	"time"

	"car_project/internal/entities"
)

// type DeliveryService struct {
// 	db *gorm.DB
// }

// func NewDeliveryService(db *gorm.DB) *DeliveryService {
// 	return &DeliveryService{db: db}
// }

// CreateTicket permet de créer un ticket de livraison (direct ou via commande)
func (s *DeliveryService) CreateTicket(ticket *entities.DeliveryTicket) error {
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if ticket.StatusID == 0 {
		ticket.StatusID = s.getPendingStatusID(tx)
	}

	now := time.Now()
	ticket.CreatedAt = now.Format("2006-10-26 12:13:15")
	ticket.UpdatedAt = now.Format("2006-10-26 12:13:15")

	if err := tx.Create(ticket).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Liste tickets visibles pour un user (admin/livreur)
func (s *DeliveryService) ListTickets(userID int, isAdmin bool, isDriver bool) ([]entities.DeliveryTicket, error) {
	var tickets []entities.DeliveryTicket
	tx := s.db

	query := tx.Preload("PickupAddress").Preload("DropoffAddress").Preload("Status")
	if !isAdmin && !isDriver {
		query = query.Where("client_id = ?", userID)
	}

	if err := query.Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}

// Update ticket (prix, status, assignation)
func (s *DeliveryService) UpdateTicket(ticketID int, update map[string]interface{}) error {
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	update["updated_at"] = time.Now()
	if err := tx.Model(&entities.DeliveryTicket{}).Where("id = ?", ticketID).Updates(update).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// func (s *DeliveryService) getPendingStatusID(tx *gorm.DB) int {
// 	var st entities.DeliveryTicketStatus
// 	if err := tx.Where("code = ?", "pending").First(&st).Error; err != nil {
// 		st = entities.DeliveryTicketStatus{
// 			Code:    "pending",
// 			Label:   "En attente",
// 			IsFinal: false,
// 		}
// 		tx.Create(&st)
// 	}
// 	return st.ID
// }
