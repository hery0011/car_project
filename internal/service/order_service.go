package service

import (
	"car_project/internal/entities"
	"time"

	"gorm.io/gorm"
)

type OrderService struct {
	db             *gorm.DB
	paymentService *PaymentService
	walletService  *WalletService
}

func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{
		db:             db,
		paymentService: &PaymentService{},
		walletService:  &WalletService{},
	}
}

// CreateOrder crée la commande et effectue le paiement si possible
func (s *OrderService) CreateOrder(
	userID int,
	pickupAddress *entities.Address,
	// dropoffAddress *entities.Address,
	items []entities.OrderItem,
	method string,
) (*entities.Order, error) {
	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// --- Création adresse de récupération si besoin ---
	if pickupAddress.AdresseID == 0 {
		pickupAddress.ClientID = userID
		if err := tx.Create(pickupAddress).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// --- Création adresse de livraison si besoin ---
	// if dropoffAddress.AdresseID == 0 {
	// 	dropoffAddress.ClientID = userID
	// 	if err := tx.Create(dropoffAddress).Error; err != nil {
	// 		tx.Rollback()
	// 		return nil, err
	// 	}
	// }

	// --- Statut initial de la commande ---
	var pendingStatus entities.OrderStatus
	if err := tx.Where("code = ?", "pending_payment").First(&pendingStatus).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	now := time.Now()
	order := &entities.Order{
		UserID:    userID,
		StatusID:  pendingStatus.ID,
		CreatedAt: now.Format("2006-01-02 15:04:05"),
		UpdatedAt: now.Format("2006-01-02 15:04:05"),
		Total:     0,
	}

	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// --- Liaison order <-> addresses ---
	if err := tx.Create(&entities.OrderAddress{
		OrderID:   order.OrderID,
		AdresseID: pickupAddress.AdresseID,
		Type:      "dropoff",
		CreatedAt: now.Format("2006-01-02 15:04:05"),
		UpdatedAt: now.Format("2006-01-02 15:04:05"),
		// DropoffAddressID: dropoffAddress.AdresseID,
	}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// --- Création des items ---
	total := 0.0
	for i := range items {
		items[i].OrderID = order.OrderID
		items[i].TotalPrice = items[i].UnitPrice * float64(items[i].Quantity)
		if err := tx.Create(&items[i]).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		total += items[i].TotalPrice
	}
	order.Total = total
	if err := tx.Save(order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// --- Traitement du paiement ---
	if err := s.handlePayment(tx, userID, order, method); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return order, nil
}

func (s *OrderService) handlePayment(tx *gorm.DB, userID int, order *entities.Order, method string) error {
	// Récupère le processor correspondant
	processor, err := s.paymentService.GetProcessor(method)
	if err != nil {
		return err
	}

	// ProcessPayment retourne désormais à la fois PaymentTransaction et Payment
	err = processor.ProcessPayment(tx, order)
	if err != nil {
		return err
	}

	// Enregistrement du paiement déjà créé par le processor (pour uniformité)
	// Si wallet, créer la transaction wallet
	if method == "wallet" {
		if err := s.walletService.DebitWallet(tx, userID, order.OrderID, order.Total); err != nil {
			return err
		}
	}

	// Met à jour le statut de la commande selon la logique du processor
	return processor.UpdateOrderStatus(tx, order)
}

func (s *OrderService) ListOrders(userID int, statusCode string) ([]entities.Order, error) {
	var orders []entities.Order
	query := s.db.Preload("Items").Preload("Status").Where("user_id = ?", userID)
	if statusCode != "" {
		query = query.Joins("JOIN order_status ON orders.status_id = order_status.id").
			Where("order_status.code = ?", statusCode)
	}
	if err := query.Order("created_at DESC").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
