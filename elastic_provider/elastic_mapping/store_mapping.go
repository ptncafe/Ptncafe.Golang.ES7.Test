package elastic_mapping

func StoreMapping() string {
	// Create a new index
	storeMapping := `
	{
		"mappings":{
			"properties":{
				"id":{
					"type": "keyword"
				},
				"name":{
					"type": "text"
				},
				"code":{
					"type": "keyword"
				},
				"shop_type": {
					"type": "keyword"
				},
				"store_level": {
					"type": "keyword"
				}
			}
		}
	}`
	return storeMapping
}