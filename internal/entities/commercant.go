package entities

type Commercant struct {
	Commercant_id int    `json:"commercant_id" gorm:"column:commercant_id"`
	Nom           string `json:"nom" gorm:"column:nom"`
	Description   string `json:"description" gorm:"column:description"`
	Adresse       string `json:"adresse" gorm:"column:adresse"`
	Telephone     string `json:"telephone" gorm:"column:telephone"`
	Email         string `json:"email" gorm:"column:email"`
}

func (Commercant) TableName() string {
	return "Commercant"
}
