package repository

import (
	"context"

	entity "ptxyz/customer-product-service/internal/domain/entity/user"
)

type UserRepository interface {
	Create(ctx context.Context, u *entity.User) error
	GetByNik(ctx context.Context, nik string) (*entity.User, error)
}