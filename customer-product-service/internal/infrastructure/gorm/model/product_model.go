package model

import (
	entity "ptxyz/customer-product-service/internal/domain/entity/product"
	"time"
)

type ProductModel struct {
	ID			uint64		`gorm:"primaryKey;unique;not null;"`
	PublicId	string		`gorm:"type:varchar(255);not null;"`
	Name		string		`gorm:"type:varchar(255);not null;"`
	Category	string		`gorm:"type:enum('White Goods','Motor','Mobil');not null"`
	HargaOtr	float64		`gorm:"type:decimal(15,2);not null"`
	CreatedAt	time.Time
	CreatedBy	int
	UpdatedAt	time.Time
	UpdatedBy	int
}

// reverse model to entity
func (m *ProductModel) ToEntity() *entity.Product {
	return &entity.Product{
		// ID: m.ID,					
		PublicId: m.PublicId,			
		Name: m.Name,				
		Category: m.Category,			
		HargaOtr: m.HargaOtr,			
		CreatedAt: m.CreatedAt,	
		CreatedBy: m.CreatedBy,	
		UpdatedAt: m.UpdatedAt,	
		UpdatedBy: m.UpdatedBy,	
	}
}

func (ProductModel) TableName() string {
	return "products"
}