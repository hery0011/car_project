package entities

type Commande struct {
	CommandeId   int     `json:"commande_id" gorm:"column:commande_id;primaryKey;autoIncrement"`
	ClientId     int     `json:"client_id" gorm:"column:client_id"`
	DateCommande string  `json:"date_commande" gorm:"column:date_commande"`
	MontantTotal float64 `json:"montant_total" gorm:"column:montant_total"`
	Status       int     `json:"status" gorm:"column:status"`
}

func (Commande) TableName() string {
	return "Commande"
}
