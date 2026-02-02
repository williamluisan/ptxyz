package auth

type LoginRequest struct {
	Nik    		string 	`json:"nik" binding:"required"`
	Password 	string  `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}