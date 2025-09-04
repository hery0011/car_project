package elastic

import (
	"car_project/internal/entities"
	"context"
	"strconv"
)

type ArticleDocument struct {
	ArticleID   int     `json:"article_id"`
	Nom         string  `json:"nom"`
	Description string  `json:"description"`
	Prix        float64 `json:"prix"`
	Stock       int     `json:"stock"`
	Categorie   string  `json:"categorie"`
	Commercant  string  `json:"commercant"`
}

func IndexArticle(article entities.ArticleResponse) error {
	doc := ArticleDocument{
		ArticleID:   article.ArticleID,
		Nom:         article.Nom,
		Description: article.Description,
		Prix:        article.Prix,
		Stock:       article.Stock,
		Categorie:   article.Categorie.Nom,
		Commercant:  article.Commercant.Nom,
	}

	_, err := GetClient().Index().
		Index("articles").
		Id(strconv.Itoa(article.ArticleID)).
		BodyJson(doc).
		Do(context.Background())

	return err
}
