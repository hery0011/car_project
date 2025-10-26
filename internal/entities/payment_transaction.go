package entities

import "time"

type PaymentTransaction struct {
	ID                   int       `gorm:"column:id;primaryKey;autoIncrement"`
	PaymentID            int       `gorm:"column:payment_id"`
	TransactionType      string    `gorm:"column:transaction_type"`
	TransactionReference string    `gorm:"column:transaction_reference"`
	Amount               float64   `gorm:"column:amount"`
	Currency             string    `gorm:"column:currency"`
	Status               string    `gorm:"column:status"`
	RawResponse          string    `gorm:"column:raw_response"`
	CreatedAt            time.Time `gorm:"column:created_at"`
}

func (PaymentTransaction) TableName() string {
	return "payment_transaction"
}
