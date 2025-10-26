package service

import (
	"car_project/internal/entities"
	"errors"
	"time"

	"gorm.io/gorm"
)

// PaymentProcessor définit les méthodes que chaque processor doit implémenter
type PaymentProcessor interface {
	ProcessPayment(tx *gorm.DB, order *entities.Order) error
	UpdateOrderStatus(tx *gorm.DB, order *entities.Order) error
}

// WalletProcessor pour paiement via wallet
type WalletProcessor struct{}

// ProcessPayment débite le wallet et crée le Payment + PaymentTransaction
func (w *WalletProcessor) ProcessPayment(tx *gorm.DB, order *entities.Order) error {
	// Récupérer le wallet de l'utilisateur
	var wallet entities.Wallet
	if err := tx.Where("user_id = ?", order.UserID).First(&wallet).Error; err != nil {
		return errors.New("wallet non trouvé")
	}

	if wallet.Balance < order.Total {
		return errors.New("solde insuffisant")
	}

	// Débiter le wallet
	wallet.Balance -= order.Total
	if err := tx.Save(&wallet).Error; err != nil {
		return err
	}

	// Créer l'enregistrement Payment
	var paymentMethod entities.PaymentMethod
	if err := tx.Where("code = ?", "wallet").First(&paymentMethod).Error; err != nil {
		return errors.New("méthode de paiement wallet introuvable")
	}

	payment := &entities.Payment{
		OrderID:         order.OrderID,
		PaymentMethodID: paymentMethod.ID,
		Amount:          order.Total,
		Status:          "completed",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if err := tx.Create(payment).Error; err != nil {
		return err
	}

	// Créer l'enregistrement PaymentTransaction
	paymentTx := &entities.PaymentTransaction{
		PaymentID:       payment.PaymentID,
		TransactionType: "capture",
		Amount:          order.Total,
		Status:          "success",
		Currency:        wallet.Currency,
		CreatedAt:       time.Now(),
	}

	if err := tx.Create(paymentTx).Error; err != nil {
		return err
	}

	// Ajouter transaction dans WalletTransaction pour journal interne
	walletTx := &entities.WalletTransaction{
		WalletID:        wallet.ID,
		TransactionType: "debit",
		Amount:          order.Total,
		Reference:       "order:" + string(order.OrderID),
		Description:     "Paiement commande via wallet",
		CreatedAt:       time.Now(),
	}
	if err := tx.Create(walletTx).Error; err != nil {
		return err
	}

	return nil
}

// UpdateOrderStatus met à jour le statut de la commande après paiement
func (w *WalletProcessor) UpdateOrderStatus(tx *gorm.DB, order *entities.Order) error {
	var status entities.OrderStatus
	if err := tx.Where("code = ?", "paid").First(&status).Error; err != nil {
		return err
	}
	return tx.Model(order).Update("status_id", status.ID).Error
}

// MobileMoneyProcessor gère les paiements externes (Mobile Money)
type MobileMoneyProcessor struct{}

// ProcessPayment crée le Payment et le PaymentTransaction avec statut pending
func (m *MobileMoneyProcessor) ProcessPayment(tx *gorm.DB, order *entities.Order) error {
	// Récupérer la méthode de paiement
	var paymentMethod entities.PaymentMethod
	if err := tx.Where("code = ?", "mobile_money").First(&paymentMethod).Error; err != nil {
		return errors.New("méthode de paiement mobile_money introuvable")
	}

	// Créer l'enregistrement Payment avec statut 'initiated' ou 'pending'
	payment := &entities.Payment{
		OrderID:         order.OrderID,
		PaymentMethodID: paymentMethod.ID,
		Amount:          order.Total,
		Status:          "initiated", // ou 'pending' selon ton choix
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if err := tx.Create(payment).Error; err != nil {
		return err
	}

	// Créer l'enregistrement PaymentTransaction
	paymentTx := &entities.PaymentTransaction{
		PaymentID:       payment.PaymentID,
		TransactionType: "capture",
		Amount:          order.Total,
		Status:          "pending",
		Currency:        "EUR", // à adapter si nécessaire
		CreatedAt:       time.Now(),
	}

	if err := tx.Create(paymentTx).Error; err != nil {
		return err
	}

	// Pas de WalletTransaction ici car c'est un paiement externe
	return nil
}

// UpdateOrderStatus ne met rien à jour directement car paiement externe
func (m *MobileMoneyProcessor) UpdateOrderStatus(tx *gorm.DB, order *entities.Order) error {
	// Le statut sera mis à jour uniquement après notification webhook / callback
	return nil
}

// Service pour récupérer processor
type PaymentService struct{}

func (p *PaymentService) GetProcessor(method string) (PaymentProcessor, error) {
	switch method {
	case "wallet":
		return &WalletProcessor{}, nil
	case "mobile_money":
		return &MobileMoneyProcessor{}, nil
	default:
		return nil, errors.New("unsupported payment method")
	}
}
