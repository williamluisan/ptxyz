package konsumertenorlimit

import (
	"context"
	entity "ptxyz/customer-product-service/internal/domain/entity/konsumer_tenor_limit"
	repository "ptxyz/customer-product-service/internal/domain/repository/konsumer_tenor_limit"
)

type KonsumerTenorLimitService interface {
	GetByPublicId(ctx context.Context, publicId string) (*entity.KonsumerTenorLimit, error)
	UpdateBalance(ctx context.Context, entity *entity.KTLUpdateBalance) error
}

type konsumerTenorLimitImpl struct {
	repo	repository.KonsumerTenorLimitRepository
}

func NewKonsumerTenorLimitService(repo repository.KonsumerTenorLimitRepository) KonsumerTenorLimitService {
	if repo == nil {
		panic("konsumer tenor limit service: repository must be implemented")
	}
	return &konsumerTenorLimitImpl{
		repo: repo,
	}
}

func (i *konsumerTenorLimitImpl) GetByPublicId(ctx context.Context, publicId string) (*entity.KonsumerTenorLimit, error) {
	return nil, nil
}

func (i *konsumerTenorLimitImpl) UpdateBalance(ctx context.Context, entityUpdate *entity.KTLUpdateBalance) error {
	return nil
}