package konsumentenorlimit

import (
	"context"
	entity "ptxyz/customer-product-service/internal/domain/entity/konsumen_tenor_limit"
)

type KonsumenTenorLimitRepository interface {
	GetByPublicId(ctx context.Context, publicId string) (*entity.KonsumenTenorLimit, error)
	UpdateBalance(ctx context.Context, entity *entity.KTLUpdateBalance) error
}