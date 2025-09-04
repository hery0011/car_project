package elastic

import (
	"context"
	"encoding/json"

	"github.com/olivere/elastic/v7"
)

func SearchArticlesAdvanced(nom, categorie, commercant string, prixMin, prixMax float64, stockMin int) ([]ArticleDocument, error) {
	boolQuery := elastic.NewBoolQuery()

	if nom != "" {
		boolQuery.Must(elastic.NewMatchQuery("nom", nom))
	}
	if categorie != "" {
		boolQuery.Filter(elastic.NewTermQuery("categorie.keyword", categorie))
	}
	if commercant != "" {
		boolQuery.Filter(elastic.NewTermQuery("commercant.keyword", commercant))
	}
	if prixMin > 0 || prixMax > 0 {
		rangeQuery := elastic.NewRangeQuery("prix")
		if prixMin > 0 {
			rangeQuery.Gte(prixMin)
		}
		if prixMax > 0 {
			rangeQuery.Lte(prixMax)
		}
		boolQuery.Filter(rangeQuery)
	}
	if stockMin > 0 {
		boolQuery.Filter(elastic.NewRangeQuery("stock").Gte(stockMin))
	}

	searchResult, err := GetClient().Search().
		Index("articles").
		Query(boolQuery).
		Do(context.Background())

	if err != nil {
		return nil, err
	}

	var articles []ArticleDocument
	for _, hit := range searchResult.Hits.Hits {
		var a ArticleDocument
		if err := json.Unmarshal(hit.Source, &a); err == nil {
			articles = append(articles, a)
		}
	}
	return articles, nil
}
