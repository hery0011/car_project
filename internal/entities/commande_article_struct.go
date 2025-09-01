package entities

type CommandeArticle struct {
	CommandeArticleId int     `json:"commande_article_id" gorm:"column:commande_article_id;primaryKey;autoIncrement"`
	CommandeId        int     `json:"commande_id" gorm:"column:commande_id"`
	ArticleId         int     `json:"article_id" gorm:"column:article_id"`
	Quantite          int     `json:"quantite" gorm:"column:quantite"`
	PrixUnitaire      float64 `json:"prix_unitaire" gorm:"column:prix_unitaire"`
}

func (CommandeArticle) TableName() string {
	return "Commande_Article"
}
