package elastic

import (
	"log"

	"github.com/olivere/elastic/v7"
)

var client *elastic.Client

func InitElasticClient(url string) {
	var err error
	client, err = elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false))
	if err != nil {
		log.Fatalf("Erreur de connexion à Elasticsearch: %v", err)
	}
}

func GetClient() *elastic.Client {
	return client
}
