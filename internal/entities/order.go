package entities

type Order struct {
	OrderID        int            `json:"order_id" gorm:"column:id;primaryKey;autoIncrement"`
	UserID         int            `json:"user_id" gorm:"column:user_id"`
	StatusID       int            `json:"status_id" gorm:"column:status_id"` // FK vers order_status
	Total          float64        `json:"total_amount" gorm:"column:total_amount"`
	CreatedAt      string         `json:"created_at" gorm:"column:created_at"`
	UpdatedAt      string         `json:"updated_at" gorm:"column:updated_at"`
	Items          []OrderItem    `json:"items" gorm:"foreignKey:OrderID;references:OrderID"`
	Status         OrderStatus    `json:"status" gorm:"foreignKey:StatusID;references:ID"`        // preload optionnel
	OrderAddresses []OrderAddress `json:"addresses" gorm:"foreignKey:OrderID;references:OrderID"` // preload pour delivery service

}

func (Order) TableName() string {
	return "orders"
}

// Payload sans user_id
type PaymentPayload struct {
	Method string `json:"paymentMethod"`
}

type OrderStatus struct {
	ID      int    `json:"id" gorm:"column:id;primaryKey"`
	Code    string `json:"code" gorm:"column:code"`
	Label   string `json:"label" gorm:"column:label"`
	IsFinal bool   `json:"is_final" gorm:"column:is_final"`
	// CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

func (OrderStatus) TableName() string {
	return "order_status"
}

type OrderAddress struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement"`
	OrderID   int    `gorm:"column:order_id"`
	AdresseID int    `gorm:"column:adresse_id"`
	Type      string `gorm:"column:type"` // 'pickup' ou 'dropoff'
	CreatedAt string `gorm:"column:created_at"`
	UpdatedAt string `gorm:"column:updated_at"`
}

// TableName retourne le nom de la table
func (OrderAddress) TableName() string {
	return "order_addresses"
}
