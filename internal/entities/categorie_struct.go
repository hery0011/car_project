package entities

type Categorie struct {
	Categorie_id int    `json:"categorie_id" gorm:"column:categorie_id"`
	Nom          string `json:"nom" gorm:"column:nom"`
	Parent_id    int    `json:"parent_id" gorm:"column:parent_id"`
}

func (Categorie) TableName() string {
	return "Categorie"
}
