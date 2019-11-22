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

func AddStore( esClient * elastic.Client, storeDto model.Store ,ctx context.Context) error {
	err := esClient.Index().Index("store").Type("store").Id(fmt.Sprint(storeDto.Id)).BodyJson(storeDto)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("AddStore Index done",)

	return nil
}
func GetStoreById (c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	var store  = elastic_provider.GetClientES().Get().Index("store").Type("store").Id(id)
	fmt.Printf("GetStoreById %s", id)
	c.JSON(200, store)
}

func UpdateStore (c *gin.Context) {
	ctx := context.Background()
	store  := model.Store{}
	if errBinding := c.ShouldBind(&store); errBinding != nil {
		c.String(http.StatusBadRequest, `the body should be formA`)
		return
	}
	if errAdd := AddStore(elastic_provider.GetClientES(), store, ctx); errAdd != nil{
		c.String(http.StatusInternalServerError, errAdd.Error())
		return
	}
	fmt.Printf("AddStore", spew.Sdump(store))
	c.JSON(200, store)
}