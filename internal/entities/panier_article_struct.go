package entities

type PanierArticle struct {
	PanierId  int `json:"panier_id" gorm:"column:panier_id"`
	ArticleId int `json:"article_id" gorm:"column:article_id"`
	Quantite  int `json:"quantite" gorm:"column:quantite"`
}

func (PanierArticle) TableName() string {
	return "Panier_Article"
}
