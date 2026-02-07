package product

import "time"

type Product struct {
	ID			uint64
	PublicId	string
	Name		string
	Category	string
	HargaOtr	float64
	CreatedAt	time.Time
	CreatedBy	int
	UpdatedAt	time.Time
	UpdatedBy	int
}