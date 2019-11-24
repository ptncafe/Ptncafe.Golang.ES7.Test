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
	"strconv"
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
	idString := c.DefaultQuery("id", "0")
	nameString := c.DefaultQuery("name", "")
	codeString := c.DefaultQuery("code", "")
	clientEs := elastic_provider.GetClientES()
	id, _ := strconv.Atoi(idString)
	ctx := context.Background()
	if id > 0 {
		var store,err  = clientEs.Get().Index("store").Id(idString).Do(ctx)
		if err != nil{
			e, ok := err.(*elastic.Error)
			if !ok {
				c.JSON(http.StatusInternalServerError, err)
			}
			c.JSON(e.Status, e)
			return
		}
		fmt.Printf("GetStoreById %s", spew.Sdump(store))
		c.JSON(200,  store.Source)
		return
	}
	if len(nameString) > 0 {
		termQuery := elastic.NewTermQuery("name", nameString)
		var store,err  = clientEs.Search().Index("store").Query(termQuery).
		From(0).Size(10).   // take documents 0-9
		Pretty(true).       // pretty print request and response JSON
		Do(ctx)
		if err != nil{
			e, ok := err.(*elastic.Error)
			if !ok {
				c.JSON(http.StatusInternalServerError, err)
			}
			c.JSON(e.Status, e)
			return
		}
		fmt.Printf("SearchStore %s", spew.Sdump(store))
		c.JSON(200,  store)
		return
	}
	if len(codeString) > 0 {
		termQuery := elastic.NewTermQuery("code", codeString)
		var store,err  = clientEs.Search().Index("store").Query(termQuery).
			From(0).Size(10).   // take documents 0-9
			Pretty(true).       // pretty print request and response JSON
			Do(ctx)
		if err != nil{
			e, ok := err.(*elastic.Error)
			if !ok {
				c.JSON(http.StatusInternalServerError, err)
			}
			c.JSON(e.Status, e)
			return
		}
		fmt.Printf("SearchStore %s", spew.Sdump(store))
		c.JSON(200,  store)
		return
	}
	c.JSON(200,  nil)
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