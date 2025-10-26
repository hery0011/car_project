package entities

import "time"

type WalletTransaction struct {
	ID              int     `json:"id" gorm:"primaryKey;autoIncrement"`
	WalletID        int     `json:"wallet_id"`
	TransactionType string  `json:"transaction_type"` // credit, debit, refund
	Amount          float64 `json:"amount"`
	Reference       string  `json:"reference"`
	Description     string
	CreatedAt       time.Time
}

func (WalletTransaction) TableName() string {
	return "wallet_transaction"
}
