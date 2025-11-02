package entities

type Articles struct {
	Article_id    int     `json:"article_id" gorm:"column:article_id;primaryKey;autoIncrement"`
	Nom           string  `json:"nom" gorm:"colum:nom"`
	Slug          string  `json:"slug" gorm:"colum:slug"`
	Description   string  `json:"description" gorm:"column:description"`
	Prix          float64 `json:"prix" gorm:"column:prix"`
	Stock         int     `json:"Stock" gorm:"column:stock"`
	Commercant_id int     `json:"commercant_id" gorm:"column:commercant_id"`
	Categorie_id  int     `json:"categorie_id" gorm:"coumn:categorie_id"`
}

func (Articles) TableName() string {
	return "article"
}

type Article struct {
	ArticleID        int     `json:"article_id" gorm:"column:article_id;primaryKey;autoIncrement"`
	Nom              string  `json:"nom" gorm:"column:nom"`
	Slug             string  `json:"slug" gorm:"column:slug"`
	ShortDescription string  `json:"short_description" gorm:"column:short_description"`
	Description      string  `json:"description" gorm:"column:description"`
	Status           string  `json:"status" gorm:"column:status;type:enum('draft','published','archived')"`
	IsActive         bool    `json:"is_active" gorm:"column:is_active"`
	CreatedAt        string  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt        string  `json:"updated_at" gorm:"column:updated_at"`
	Prix             float64 `json:"prix" gorm:"column:prix"`
	Stock            int     `json:"stock" gorm:"column:stock"`
	CommercantID     int     `json:"commercant_id" gorm:"column:commercant_id"`

	// Relations
	Images     []ArticleImage `json:"images" gorm:"foreignKey:Article_id"`
	Categories []Categorie    `json:"categories" gorm:"many2many:article_category;joinForeignKey:article_id;joinReferences:categorie_id"`
	Commercant Commercant     `json:"commercant" gorm:"foreignKey:Commercant_id"`
}

type ArticleResponse struct {
	ArticleID   int            `json:"id"`
	Nom         string         `json:"nom"`
	Slug        string         `json:"slug"`
	Description string         `json:"description"`
	Prix        float64        `json:"prix"`
	Stock       int            `json:"stock"`
	Categorie   []Categorie    `json:"categorie"`
	Commercant  Commercant     `json:"commercant"`
	Images      []ArticleImage `json:"images"`
}

func (Article) TableName() string {
	return "article"
}

// type PersonAddress struct {
// 	ArticleID    int `gorm:"primaryKey"`
// 	Categorie_id int `gorm:"primaryKey"`
// }

// ImagePayload représente les données et métadonnées d'une seule image envoyée
type ImagePayload struct {
	Base64Data string `json:"base64_data" binding:"required"` // L'image encodée
	Largeur    int    `json:"largeur"`
	Hauteur    int    `json:"hauteur"`
	Ordre      int    `json:"ordre"`
	Type       string `json:"type"` // 'main', 'gallery', 'thumbnail', etc.
	Taille     string `json:"taille"`
}

// ArticleCreateRequest structure pour la création d'un article avec plusieurs images
type ArticleCreateRequest struct {
	Nom          string         `json:"nom" binding:"required"`
	Description  string         `json:"description" binding:"required"`
	Prix         float64        `json:"prix" binding:"required,gt=0"`
	Stock        int            `json:"stock" binding:"gte=0"`
	CommercantID int            `json:"commercant_id" binding:"required"`
	CategorieIDs []int          `json:"categorie_ids" binding:"required"` // Changement pour gérer plusieurs catégories
	Images       []ImagePayload `json:"images" binding:"required"`        // 🔑 Changement: Slice de données d'image
}

// --------------------
// Struct pour l'update
// --------------------
type ArticleUpdateRequest struct {
	ArticleID    uint                `json:"id" binding:"required"`
	Nom          string              `json:"nom" binding:"required"`
	Description  string              `json:"description"`
	Prix         float64             `json:"prix"`
	Stock        int                 `json:"stock"`
	CommercantID uint                `json:"commercant_id,omitempty"` // optionnel si pas modifiable
	CategorieIDs []uint              `json:"categorie_ids"`
	Images       []ArticleImageInput `json:"images"`
}

type ArticleImageInput struct {
	ImageID    uint   `json:"image_id,omitempty"` // ID existante si modification/suppression
	ToDelete   bool   `json:"to_delete,omitempty"`
	Base64Data string `json:"base64_data,omitempty"` // pour nouvelle image
	Ordre      int    `json:"ordre,omitempty"`
	Type       string `json:"type,omitempty"` // main, gallery, thumbnail
	Taille     string `json:"taille,omitempty"`
	Largeur    int    `json:"largeur,omitempty"`
	Hauteur    int    `json:"hauteur,omitempty"`
}

type ArticleFilterRequest struct {
	Category     string `json:"category"`
	ProductText  string `json:"productText"`
	MerchantText string `json:"merchantText"`
	Prix         struct {
		Lower float64 `json:"lower"`
		Upper float64 `json:"upper"`
	} `json:"prix"`
}
