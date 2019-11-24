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
				}
			}
		}
	}`
	return storeMapping
}