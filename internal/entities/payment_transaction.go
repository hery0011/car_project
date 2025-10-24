package entities

import "time"

type PaymentTransaction struct {
	ID              int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	OrderID         int       `json:"order_id" gorm:"column:order_id"`
	Method          string    `json:"method" gorm:"column:method"` // wallet, mobile_money
	Amount          float64   `json:"amount" gorm:"column:amount"`
	Status          string    `json:"status" gorm:"column:status"` // pending, paid, failed
	TransactionID   string    `json:"transaction_id" gorm:"column:transaction_id"`
	CreatedAt       time.Time `json:"created_at" gorm:"column:created_at"`
	TransactionType string    `json:"transaction_type" gorm:"column:transaction_type"` // authorization, capture, refund, void
}

func (PaymentTransaction) TableName() string {
	return "payment_transaction"
}
