package konsumertenorlimit

import (
	"context"
	entity "ptxyz/customer-product-service/internal/domain/entity/konsumer_tenor_limit"
)

type KonsumerTenorLimitRepository interface {
	GetByPublicId(ctx context.Context, publicId string) (*entity.KonsumerTenorLimit, error)
	UpdateBalance(ctx context.Context, entity *entity.KTLUpdateBalance) error
}