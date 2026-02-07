package model

import (
	entity "ptxyz/customer-product-service/internal/domain/entity/konsumentenorlimit"
	"time"
)

type KonsumenTenorLimitModel struct {
	ID			uint64		`gorm:"primaryKey;unique;not null;"`
	PublicId	string		`gorm:"type:varchar(255);not null;`
	UserId		int			`gorm:"type:bigint unsigned;not null;`
	TenorRefId	int			`gorm:"type:bigint unsigned;not null;`
	Limit		float64		`gorm:"type:decimal(15,0);"`
	UsedAmount	float64		`gorm:"type:decimal(15,0);"`
	Balance 	float64		`gorm:"type:decimal(15,0);`
	CreatedAt	time.Time
	CreatedBy	int
	UpdatedAt	time.Time
	UpdatedBy	int
}

func (KonsumenTenorLimitModel) TableName() string {
	return "konsumen_tenor_limit"
}

// reverse model to entity
func (m *KonsumenTenorLimitModel) ToEntity() *entity.KonsumenTenorLimit {
	return &entity.KonsumenTenorLimit{
		PublicId: m.PublicId,
		UserId: m.UserId,
		TenorRefId: m.TenorRefId,
		Limit: m.Limit,
		UsedAmount: m.UsedAmount,
		Balance: m.Balance,
		CreatedAt: m.CreatedAt,
		CreatedBy: m.CreatedBy,
		UpdatedAt: m.UpdatedAt,
		UpdatedBy: m.UpdatedBy,
	}
}