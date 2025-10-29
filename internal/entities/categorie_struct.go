package entities

type Categorie struct {
	Categorie_id  int         `json:"categorie_id" gorm:"column:categorie_id;primaryKey"`
	Nom           string      `json:"nom" gorm:"column:nom"`
	Parent_id     int         `json:"parent_id" gorm:"column:parent_id"`
	ImageUrl      string      `json:"image_url" gorm:"column:image_url"`
	SubCategories []Categorie `json:"images" gorm:"foreignKey:parent_id"`
}

type CategoryResponse struct {
	CategoryId    uint                `json:"categoryId"`
	Nom           string              `json:"nom"`
	ImageUrl      string              `json:"imageUrl,omitempty"`
	SubCategories []*CategoryResponse `json:"subCategories,omitempty"` // pointeurs
}

func (Categorie) TableName() string {
	return "categorie"
}
