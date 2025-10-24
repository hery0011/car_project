package entities

type WalletTransaction struct {
	WalletTransactionID int     `json:"wallet_transaction_id" gorm:"primaryKey;autoIncrement"`
	WalletID            int     `json:"wallet_id" gorm:"column:wallet_id"`
	TransactionType     string  `json:"transaction_type" gorm:"column:transaction_type"`
	Amount              float64 `json:"amount" gorm:"column:amount"`
	Reference           string  `json:"reference" gorm:"column:reference"`
	Description         string  `json:"description" gorm:"column:description"`
}

func (WalletTransaction) TableName() string {
	return "wallet_transaction"
}
