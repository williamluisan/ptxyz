package product

import (
	"context"
	entity "ptxyz/customer-product-service/internal/domain/entity/product"
	repository "ptxyz/customer-product-service/internal/domain/repository/product"
)

type ProductService interface {
	GetByPublicId(ctx context.Context, publicId string) (*entity.Product, error)
}

type productServiceImpl struct {
	repo	repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	if repo == nil {
		panic("product service: repository must be implemented.")
	}
	return &productServiceImpl{
		repo: repo,
	}
}

func (i *productServiceImpl) GetByPublicId(ctx context.Context, publicId string) (*entity.Product, error) {
	data, err := i.repo.GetByPublicId(ctx, publicId)
	if err != nil {
		return nil, err
	}

	return data, nil
}