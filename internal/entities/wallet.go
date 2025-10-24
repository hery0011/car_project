package entities

type Wallet struct {
	WalletID int     `json:"wallet_id" gorm:"column:wallet_id;primaryKey;autoIncrement"`
	UserID   int     `json:"user_id" gorm:"column:user_id"`
	Balance  float64 `json:"balance" gorm:"column:balance"`
}

func (Wallet) TableName() string {
	return "wallet"
}
