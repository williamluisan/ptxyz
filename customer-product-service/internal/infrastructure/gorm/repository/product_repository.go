package repository

import (
	"context"

	"gorm.io/gorm"

	entity "ptxyz/customer-product-service/internal/domain/entity/product"
	repository "ptxyz/customer-product-service/internal/domain/repository/product"
	gormModel "ptxyz/customer-product-service/internal/infrastructure/gorm/model"
)

type productRepository struct {
	db *gorm.DB
}

func NewProdctRepository(db *gorm.DB) repository.ProductRepository {
	return &productRepository{db: db}
}

var mProductModel gormModel.ProductModel

func (r *productRepository) GetByPublicId(ctx context.Context, publicId string) (*entity.Product, error) {
	product, err := gorm.G[gormModel.ProductModel](r.db).Where("public_id = ?", publicId).First(ctx);
	if err != nil {
		return nil, err
	}

	return product.ToEntity(), nil
}