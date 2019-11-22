package model

type Store struct {
	Id     int       `json:"id"`
	Name  string    `json:"name"`
	Code  string 	`json:"code"`
	ShopType int `json:"shop_type"`
	StoreLevel int `json:"store_level"`
}