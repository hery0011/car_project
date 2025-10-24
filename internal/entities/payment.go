package entities

type Payment struct {
	PaymentID       int                  `json:"payment_id" gorm:"column:payment_id;primaryKey;autoIncrement"`
	OrderID         int                  `json:"order_id" gorm:"column:order_id"`
	PaymentMethodID int                  `json:"payment_method_id" gorm:"column:payment_method_id"`
	Amount          float64              `json:"amount" gorm:"column:amount"`
	Status          string               `json:"status" gorm:"column:status"`
	Transactions    []PaymentTransaction `json:"transactions" gorm:"foreignKey:PaymentID;references:PaymentID"`
}

func (Payment) TableName() string {
	return "payment"
}
