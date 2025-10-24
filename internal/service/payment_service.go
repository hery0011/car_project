package service

import (
	"car_project/internal/entities"
	"errors"
	"time"
)

// PaymentProcessor interface pour chaque méthode de paiement
type PaymentProcessor interface {
	ProcessPayment(orderID int, amount float64) (*entities.PaymentTransaction, error)
}

// WalletProcessor
type WalletProcessor struct{}

func (w *WalletProcessor) ProcessPayment(orderID int, amount float64) (*entities.PaymentTransaction, error) {
	// Débit wallet ici (logique métier réelle)
	return &entities.PaymentTransaction{
		OrderID:         orderID,
		Method:          "wallet",
		Amount:          amount,
		Status:          "paid",
		TransactionType: "capture",
		CreatedAt:       time.Now(),
	}, nil
}

// MobileMoneyProcessor
type MobileMoneyProcessor struct{}

func (m *MobileMoneyProcessor) ProcessPayment(orderID int, amount float64) (*entities.PaymentTransaction, error) {
	// Logique MobileMoney réelle ici
	return &entities.PaymentTransaction{
		OrderID:         orderID,
		Method:          "mobile_money",
		Amount:          amount,
		Status:          "paid",
		TransactionType: "capture",
		CreatedAt:       time.Now(),
	}, nil
}

// PaymentService permet de récupérer le processor
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
