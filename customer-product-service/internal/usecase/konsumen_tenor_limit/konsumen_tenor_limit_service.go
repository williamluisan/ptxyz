package konsumentenorlimit

import (
	"context"
	entity "ptxyz/customer-product-service/internal/domain/entity/konsumen_tenor_limit"
	repository "ptxyz/customer-product-service/internal/domain/repository/konsumen_tenor_limit"
)

type KonsumenTenorLimitService interface {
	GetByPublicId(ctx context.Context, publicId string) (*entity.KonsumenTenorLimit, error)
	UpdateBalance(ctx context.Context, entity *entity.KTLUpdateBalance) error
}

type konsumenTenorLimitImpl struct {
	repo	repository.KonsumenTenorLimitRepository
}

func NewKonsumenTenorLimitService(repo repository.KonsumenTenorLimitRepository) KonsumenTenorLimitService {
	if repo == nil {
		panic("konsumen tenor limit service: repository must be implemented")
	}
	return &konsumenTenorLimitImpl{
		repo: repo,
	}
}

func (i *konsumenTenorLimitImpl) GetByPublicId(ctx context.Context, publicId string) (*entity.KonsumenTenorLimit, error) {
	data, err := i.repo.GetByPublicId(ctx, publicId)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (i *konsumenTenorLimitImpl) UpdateBalance(ctx context.Context, entityUpdate *entity.KTLUpdateBalance) error {
	if err := i.repo.UpdateBalance(ctx, entityUpdate); err != nil {
		return err
	}

	return nil
}