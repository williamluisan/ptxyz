package auth

type VerifyRequest struct {
	Nik		    string	`json:"nik" binding:"required"`
	Password 	string	`json:"password" binding:"required"`
}

type VerifyResponse struct {
	UserPublicId string `json:"user_public_id"`
}
