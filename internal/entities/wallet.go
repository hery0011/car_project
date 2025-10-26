package entities

type Wallet struct {
	ID        int     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    int     `json:"user_id"`
	Balance   float64 `json:"balance"`
	Currency  string  `json:"currency"`
	UpdatedAt string  `json:"updated_at"`
}

func (Wallet) TableName() string {
	return "wallet"
}
