package entities

import "time"

type Payment struct {
	PaymentID       int                  `json:"payment_id" gorm:"column:id;primaryKey;autoIncrement"`
	OrderID         int                  `json:"order_id" gorm:"column:order_id"`
	PaymentMethodID int                  `json:"payment_method_id" gorm:"column:method_id"`
	Amount          float64              `json:"amount" gorm:"column:amount"`
	Status          string               `json:"status" gorm:"column:status"`
	Transactions    []PaymentTransaction `json:"transactions" gorm:"foreignKey:PaymentID;references:PaymentID"` // attention aux noms ci-dessous
	CreatedAt       time.Time            `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       time.Time            `json:"updated_at" gorm:"column:updated_at"`
}

func (Payment) TableName() string {
	return "payments"
}

type PaymentMethod struct {
	ID          int    `gorm:"column:id;primaryKey"`
	Code        string `gorm:"column:code"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	IsActive    bool   `gorm:"column:is_active"`
	CreatedAt   string `gorm:"column:created_at"`
}

func (PaymentMethod) TableName() string {
	return "payment_method"
}

// func GetPaymentMethod(code string) *PaymentMethod {
// 	data := map[string]PaymentMethod{
// 		"stripe": {
// 			ID:          1,
// 			Code:        "stripe",
// 			Name:        "Paiement carte via Stripe",
// 			Description: "",
// 			IsActive:    true,
// 			CreatedAt:   time.Date(2025, 10, 24, 16, 48, 28, 0, time.UTC),
// 		},
// 		"paypal": {
// 			ID:          2,
// 			Code:        "paypal",
// 			Name:        "Paiement PayPal",
// 			Description: "",
// 			IsActive:    true,
// 			CreatedAt:   time.Date(2025, 10, 24, 16, 48, 28, 0, time.UTC),
// 		},
// 		"cash": {
// 			ID:          3,
// 			Code:        "cash",
// 			Name:        "Paiement à la livraison",
// 			Description: "",
// 			IsActive:    true,
// 			CreatedAt:   time.Date(2025, 10, 24, 16, 48, 28, 0, time.UTC),
// 		},
// 		"wallet": {
// 			ID:          4,
// 			Code:        "wallet",
// 			Name:        "Portefeuille interne",
// 			Description: "",
// 			IsActive:    true,
// 			CreatedAt:   time.Date(2025, 10, 24, 16, 48, 28, 0, time.UTC),
// 		},
// 	}

// 	// Retourne nil si code non trouvé
// 	if val, ok := data[code]; ok {
// 		return &val
// 	}

// 	return nil
// }
