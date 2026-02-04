package konsumentenorlimit

import "time"

type KonsumenTenorLimit struct {
	ID			uint64
	PublicId	string
	UserId		int
	TenorRefId	int
	Limit		float64
	UsedAmount	float64
	Balance 	float64
	CreatedAt	time.Time
	CreatedBy	int
	UpdatedAt	time.Time
	UpdatedBy	int
}

// DTO
type KTLUpdateBalance struct {
	ID			uint64
	UsedAmount	float64
	Balance 	float64
	UpdatedBy	int
}