package entities

type ArticleImage struct {
	Image_id   int    `json:"image_id" gorm:"column:image_id"`
	Article_id int    `json:"article_id" gorm:"column:article_id"`
	Url        string `json:"url" gorm:"column:url"`
	Largeur    int    `json:"largeur" gorm:"column:largeur"`
	Hauteur    int    `json:"hauteur" gorm:"column:hauteur"`
	Ordre      int    `json:"order" gorm:"column:ordre"`
	Type       string `json:"type" gorm:"column:type; type:enum('main', 'gallery', 'thumbnail')"`
	Taille     string `json:"taille" gorm:"column:taille"`
}

func (ArticleImage) TableName() string {
	return "article_image"
}
