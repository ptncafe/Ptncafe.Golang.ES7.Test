package main

import (
	"Ptncafe.Golang.ES7.Test/elastic_provider"
	"Ptncafe.Golang.ES7.Test/model"
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"net/http"
)

func main() {
	ctx := context.Background()
	elastic_provider.InitES(ctx)
	//var storeDto =  model.Store{
	//	Code: "Code_Ngin_Shop",
	//	Name: "Name Nguyên Phạm",
	//	Id: 1,
	//	ShopType: 1,
	//	StoreLevel: 1,
	//}
	//
	//AddStore(elastic_provider.GetClientES(), storeDto, ctx)
	// Setting up Gin
	r := gin.Default()
	//r.POST("/query", graphqlHandler())
	//r.GET("/", playgroundHandler())
	r.GET("/store",GetStoreById)
	r.POST("/store",UpdateStore)

	r.Run()


}

func AddStore( esClient  elastic.Client, storeDto model.Store ,ctx context.Context) error {
	store, err := esClient.Index().Index("store").Id(fmt.Sprint(storeDto.Id)).BodyJson(storeDto).Do(ctx)
	if err != nil {
		// Handle error
		return err
	}
	fmt.Printf("AddStore Index done %s",spew.Sdump(store))

	return nil
}
func GetStoreById (c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	clientEs := elastic_provider.GetClientES()
	ctx := context.Background()

	var store,err  = clientEs.Get().Index("store").Id(id).Do(ctx)
	if err!=nil{
		e, ok := err.(*elastic.Error)
		if !ok {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(e.Status, e)
		return
	}
	fmt.Printf("GetStoreById %s", spew.Sdump(store))
	c.JSON(200,  store.Source)
}

func UpdateStore (c *gin.Context) {
	ctx := context.Background()
	store  := model.Store{}
	if errBinding := c.ShouldBind(&store); errBinding != nil {
		c.String(http.StatusBadRequest, `the body should be formA`)
		return
	}

	clientEs := elastic_provider.GetClientES()

	if errAdd := AddStore(clientEs, store, ctx); errAdd != nil{
		e, ok := errAdd.(*elastic.Error)
		if !ok {
			c.JSON(http.StatusInternalServerError, errAdd)

		}
		c.JSON(e.Status, e)
		return
	}
	fmt.Printf("AddStore", spew.Sdump(store))
	c.JSON(200, store)
}