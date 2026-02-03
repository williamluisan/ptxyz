package product

import (
	"context"
	entity "ptxyz/customer-product-service/internal/domain/entity/product"
)

type ProductRepository interface {
	GetByPublicId(ctx context.Context, publicId string) (*entity.Product, error)
}