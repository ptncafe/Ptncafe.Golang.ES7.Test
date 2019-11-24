package elastic_mapping

func StoreMapping() string {
	// Create a new index
	storeMapping := `
	{
		"mappings":{
			"properties":{
				"id":{
					"type": "long"
				},
				"name":{
					"type": "text"
				},
				"code":{
					"type":  "keyword"
				},
				"shop_type":{
					"type": "integer",
					"null_value": 0
				},
				"store_level":{
					"type": "integer",
					"null_value": 0
				}
			}
		}
	}`
	return storeMapping
}