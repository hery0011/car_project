package entities

type OrderItem struct {
	OrderItemID int     `json:"order_item_id" gorm:"column:id;primaryKey;autoIncrement"`
	OrderID     int     `json:"order_id" gorm:"column:order_id"`
	ArticleID   int     `json:"article_id" gorm:"column:article_id"`
	Quantity    int     `json:"quantity" gorm:"column:quantity"`
	UnitPrice   float64 `json:"price" gorm:"column:unit_price"`
	TotalPrice  float64 `json:"total_price" gorm:"column:total_price"`
}

func (OrderItem) TableName() string {
	return "order_items"
}
