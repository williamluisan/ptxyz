package konsumentenorlimit

type UpdateKTLBalanceRequest struct {
	PublicId 	string	`json:"public_id" binding:"required"`
	UsedAmount	float64	`json:"used_amount" binding:"required,gt=0"`
	Balance 	float64 `json:"balance" binding:"required,gte=0"`
}