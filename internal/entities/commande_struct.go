package entities

type Commande struct {
	CommandeId    int     `json:"commande_id" gorm:"column:commande_id;primaryKey;autoIncrement"`
	ClientId      int     `json:"client_id" gorm:"column:client_id"`
	DateCommande  string  `json:"date_commande" gorm:"column:date_commande"`
	MontantTotal  float64 `json:"montant_total" gorm:"column:montant_total"`
	StatusId      int     `json:"status_id" gorm:"column:status_id"`
	LivreurAssign int     `json:"livreur_assign" gorm:"column:livreur_assign"`
}

func (Commande) TableName() string {
	return "Commande"
}
