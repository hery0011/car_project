package entities

type Panier struct {
	PanierId     int    `json:"panier_id" gorm:"column:panier_id;primaryKey;autoIncrement"`
	ClientId     int    `json:"client_id" gorm:"column:client_id"`
	DateCreation string `json:"date_creation" gorm:"column:date_creation"`
	Status_id    int    `json:"status_id" gorm:"column:status_id"`
}

func (Panier) TableName() string {
	return "Panier"
}

type PayloadPanier struct {
	ClientId  int `json:"client_id"`
	ArticleId int `json:"article_id"`
	Quantite  int `json:"quantite"`
}
