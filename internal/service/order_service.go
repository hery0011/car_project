package service

import (
	"car_project/internal/entities"
	"time"

	"gorm.io/gorm"
)

type OrderService struct {
	db             *gorm.DB
	paymentService *PaymentService
}

func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{
		db:             db,
		paymentService: &PaymentService{},
	}
}

// CreateOrder gère la création d'une commande complète
func (s *OrderService) CreateOrder(userID int, address *entities.Address, items []entities.OrderItem, method string) (*entities.Order, error) {
	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Création adresse si inexistante
	if address.AdresseID == 0 {
		address.ClientID = userID
		if err := tx.Create(address).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	order := &entities.Order{
		UserID:    userID,
		AddressID: address.AdresseID,
		Status:    "pending",
		CreatedAt: time.Now(),
	}
	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	total := 0.0
	// merchants := map[int]bool{}
	for i := range items {
		items[i].OrderID = order.OrderID
		if err := tx.Create(&items[i]).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		total += items[i].UnitPrice * float64(items[i].Quantity)
		// merchants[items[i].MerchantID] = true
	}

	// Paiement via PaymentService
	processor, err := s.paymentService.GetProcessor(method)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	paymentRecord, err := processor.ProcessPayment(order.OrderID, total)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Create(paymentRecord).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(order).Update("status", "paid").Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Retourne order et liste des merchants pour notification
	// order.MerchantIDs = merchants
	return order, nil
}
