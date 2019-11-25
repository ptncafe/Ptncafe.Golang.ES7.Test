package elastic_provider
import (
	"Ptncafe.Golang.ES7.Test/elastic_provider/elastic_mapping"
	"context"
	_ "context"
	"github.com/olivere/elastic/v7"
	"log"
	"net/http"
	"os"
	"time"
)

var clientEs  elastic.Client

func InitES(ctx context.Context) (* elastic.Client, error){
	client, err := elastic.NewClient(
		elastic.SetURL("https://u1csqf6a0i:ujcu1051my@ptncafe-test-golang-8394376725.ap-southeast-2.bonsaisearch.net:443"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		//elastic.SetRetrier(NewCustomRetrier()),
		elastic.SetGzip(true),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		elastic.SetHeaders(http.Header{
			"X-Caller-Id": []string{"..."},
		}),
	)
	if err != nil {
		panic(err.Error())
	}
	err = elastic_mapping.ElasticMapping(client, ctx)
	if err != nil {
		panic(err.Error())
	}
	clientEs = *client
	return client,err
}


func GetClientES() elastic.Client{
	return clientEs
}
