package konsumertenorlimit

type KonsumerTenorLimit struct {
	ID			uint64
	PublicId	string
	UserId		int
	TenorRefId	int
	Limit		float64
	UsedAmount	float64
	Balance 	float64
}

// DTO
type KTLUpdateBalance struct {
	ID			uint64
	Limit 		float64
	UsedAmount	float64
	Balance 	float64
}