package elastic_mapping
import (
	"context"
	"github.com/olivere/elastic/v7"
)


func ElasticMapping( esClient *elastic.Client, ctx context.Context ) error {
	// Check if the index called "store" exists
	exists, err := esClient.IndexExists("store").Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		createIndex, err := esClient.CreateIndex("store").BodyString(StoreMapping()).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}
	return err
}