package repository

import (
	"context"

	entity "ptxyz/customer-product-service/internal/domain/entity/user"
	repository "ptxyz/customer-product-service/internal/domain/repository/user"
	gormModel "ptxyz/customer-product-service/internal/infrastructure/gorm/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

var mUserModel gormModel.UserModel

func (r *userRepository) GetByNik(ctx context.Context, nik string) (*entity.User, error) {
	user, err := gorm.G[gormModel.UserModel](r.db).Where("nik = ?", nik).First(ctx);
	if err != nil {
		return nil, err
	}

	return user.ToEntity(), nil
}

func (r *userRepository) Create(ctx context.Context, u *entity.User) error {
	data := mUserModel.FromEntity(u)

	if err := gorm.G[gormModel.UserModel](r.db).Create(ctx, data); err != nil {
		return err
	}

	return nil
}