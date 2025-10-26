package entities

type DeliveryTicket struct {
	ID                 int      `gorm:"column:id;primaryKey;autoIncrement"`
	NomTicket          string   `gorm:"column:nom_ticket"`
	OrderID            int      `gorm:"column:order_id"`
	ClientID           int      `gorm:"column:client_id"`
	PickupAddressID    int      `gorm:"column:pickup_address_id"`
	DropoffAddressID   int      `gorm:"column:dropoff_address_id"`
	DeliveryPrice      *float64 `gorm:"column:delivery_price"` // pointer allows NULL
	PriceLastUpdatedBy *int     `gorm:"column:price_last_updated_by"`
	StatusID           int      `gorm:"column:status_id"`
	AssignedTo         *int     `gorm:"column:assigned_to"`
	CreatedAt          string   `gorm:"column:created_at"`
	UpdatedAt          string   `gorm:"column:updated_at"`

	PickupAddress  Address              `gorm:"foreignKey:PickupAddressID;references:AdresseID"`
	DropoffAddress Address              `gorm:"foreignKey:DropoffAddressID;references:AdresseID"`
	Status         DeliveryTicketStatus `gorm:"foreignKey:StatusID;references:ID"` // <-- relation ajoutée

}

func (DeliveryTicket) TableName() string {
	return "delivery_tickets"
}

// DeliveryTicketStatus représente le statut d'un ticket de livraison
type DeliveryTicketStatus struct {
	ID        int    `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Code      string `json:"code" gorm:"column:code;unique;not null"` // 'pending', 'assigned', 'picked', 'delivered', 'cancelled'
	Label     string `json:"label" gorm:"column:label;not null"`
	IsFinal   bool   `json:"is_final" gorm:"column:is_final;default:false"`
	CreatedAt string `json:"created_at" gorm:"column:created_at;autoCreateTime"`
}

func (DeliveryTicketStatus) TableName() string {
	return "delivery_ticket_status"
}
