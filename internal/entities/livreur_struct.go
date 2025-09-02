package entities

type Livreur struct {
	Livreur_id    int    `json:"livreur_id" gorm:"column:livreur_id;primaryKey;autoIncrement"`
	Nom           string `json:"nom" gorm:"column:nom"`
	Telephone     string `json:"telephone" gorm:"column:telephone"`
	Vehicule      string `json:"vehicule" gorm:"column:vehicule"`
	ZoneLivraison string `json:"zone_livraison" gorm:"column:zone_livraison"`
}

func (Livreur) TableName() string {
	return "Livreur"
}
