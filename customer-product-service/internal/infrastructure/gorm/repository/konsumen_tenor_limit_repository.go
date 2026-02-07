package repository

import (
	"context"

	"gorm.io/gorm"

	entity "ptxyz/customer-product-service/internal/domain/entity/konsumentenorlimit"
	repository "ptxyz/customer-product-service/internal/domain/repository/konsumentenorlimit"
	gormModel "ptxyz/customer-product-service/internal/infrastructure/gorm/model"
)

type konsumenTenorLimitRepository struct {
	db *gorm.DB
}

func NewKonsumenTenorLimitRepository(db *gorm.DB) repository.KonsumenTenorLimitRepository {
	return &konsumenTenorLimitRepository{db: db}
}

var mKTLModel gormModel.KonsumenTenorLimitModel

func (r *konsumenTenorLimitRepository) GetByPublicId(ctx context.Context, publicId string) (*entity.KonsumenTenorLimit, error) {
	data, err := gorm.G[gormModel.KonsumenTenorLimitModel](r.db).Where("public_id = ?", publicId).First(ctx);
	if err != nil {
		return nil, err
	}

	return data.ToEntity(), nil
}

func (r *konsumenTenorLimitRepository) UpdateBalance(ctx context.Context, entity *entity.KTLUpdateBalance) error {
	_, err := gorm.G[gormModel.KonsumenTenorLimitModel](r.db.Debug()).Where("public_id = ?", entity.PublicId).
		Updates(ctx, gormModel.KonsumenTenorLimitModel{
			UsedAmount: entity.UsedAmount,
			Balance: entity.Balance,
			UpdatedBy: 1,
		})

	return err
}