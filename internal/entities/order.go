package entities

import "time"

type Order struct {
	OrderID   int         `json:"order_id" gorm:"column:order_id;primaryKey;autoIncrement"`
	UserID    int         `json:"user_id" gorm:"column:user_id"`
	AddressID int         `json:"address_id" gorm:"column:address_id"`
	Total     float64     `json:"total" gorm:"column:total"`
	Status    string      `json:"status" gorm:"column:status"` // pending, paid
	CreatedAt time.Time   `json:"created_at" gorm:"column:created_at"`
	Items     []OrderItem `json:"items" gorm:"foreignKey:OrderID;references:OrderID"`
}

func (Order) TableName() string {
	return "order"
}
