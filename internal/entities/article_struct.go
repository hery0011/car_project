package entities

type Articles struct {
	Article_id    int     `json:"article_id" gorm:"column:article_id;primaryKey;autoIncrement"`
	Nom           string  `json:"nom" gorm:"colum:nom"`
	Description   string  `json:"description" gorm:"column:description"`
	Prix          float64 `json:"prix" gorm:"column:prix"`
	Stock         int     `json:"Stock" gorm:"column:stock"`
	Commercant_id int     `json:"commercant_id" gorm:"column:commercant_id"`
	Categorie_id  int     `json:"categorie_id" gorm:"coumn:categorie_id"`
}

func (Article) TableName() string {
	return "Article"
}

type Article struct {
	Article_id    int     `json:"article_id" gorm:"column:article_id;primaryKey;autoIncrement"`
	Nom           string  `json:"nom" gorm:"column:nom"`
	Description   string  `json:"description" gorm:"column:description"`
	Prix          float64 `json:"prix" gorm:"column:prix"`
	Stock         int     `json:"stock" gorm:"column:stock"`
	Commercant_id int     `json:"commercant_id" gorm:"column:commercant_id"`
	Categorie_id  int     `json:"categorie_id" gorm:"column:categorie_id"`

	// Relations
	Images     []ArticleImage `json:"images" gorm:"foreignKey:Article_id;references:Article_id"`
	Categorie  Categorie      `json:"categorie" gorm:"foreignKey:Categorie_id;references:Categorie_id"`
	Commercant Commercant     `json:"commercant" gorm:"foreignKey:Commercant_id;references:Commercant_id"`
}

type ArticleResponse struct {
	ArticleID   int            `json:"article_id"`
	Nom         string         `json:"nom"`
	Description string         `json:"description"`
	Prix        float64        `json:"prix"`
	Stock       int            `json:"stock"`
	Categorie   Categorie      `json:"categorie"`
	Commercant  Commercant     `json:"commercant"`
	Images      []ArticleImage `json:"images"`
}
