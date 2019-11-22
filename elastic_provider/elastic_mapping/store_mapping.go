package elastic_mapping

func StoreMapping() string {
	// Create a new index
	storeMapping := `
	{
		"mappings":{
			"properties":{
				"name":{
					"type": "text"
				},
				"code":{
					"type": "keyword"
				}
			
			}
		}
	}`
	return storeMapping
}