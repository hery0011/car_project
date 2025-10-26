package entities

type Address struct {
	AdresseID  int     `json:"adresse_id" gorm:"column:adresse_id;primaryKey;autoIncrement"`
	ClientID   int     `json:"client_id" gorm:"column:client_id"` // FK vers User
	Rue        string  `json:"rue" gorm:"column:rue"`
	Ville      string  `json:"ville" gorm:"column:ville"`
	CodePostal string  `json:"code_postal" gorm:"column:code_postal"`
	Pays       string  `json:"pays" gorm:"column:pays"`
	Latitude   float64 `json:"latitude" gorm:"column:latitude"`
	Longitude  float64 `json:"longitude" gorm:"column:longitude;not null"`
}

// TableName indique à GORM la table correspondante
func (Address) TableName() string {
	return "adresse"
}
