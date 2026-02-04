package client

type getProductResponse struct {
	Success bool `json:"success"`
	Data struct {
		PublicId 	string  	`json:"public_id"`
		Name     	string  	`json:"name"`
		Category	string		`json:"category"`
		HargaOtr    float64 	`json:"harga_otr"`
	} `json:"data"`
}