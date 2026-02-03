package model

import (
	entity "ptxyz/transaction-service/internal/domain/entity/transaction"
	"time"
)

type TransactionModel struct {
	ID							uint64		`gorm:"primaryKey;unique;not null"`
	PublicId					string		`gorm:"type:varchar(255); not null;"`
	KonsumenTenorLimitPublicId	string		`gorm:"type:varchar(255); not null;"`
	ProductPublicId				string		`gorm:"type:varchar(255); not null;`
	NomorKontrak				string		`gorm:"type:varchar(100); not null;`
	HargaOtr					string		`gorm:"type:decimal(15,0); not null;`
	AdminFee					string		`gorm:"type:decimal(15,0)"`
	JumlahCicilan				string		`gorm:"type:int;not null"`
	JumlahBunga					string		`gorm:"type:int"`
	NamaAset					string		`gorm:"not null;"`
	CreatedAt					time.Time
	CreatedBy					int
	UpdatedAt					time.Time
	UpdatedBy					int
}

func (TransactionModel) TableName() string {
	return "transactions"
}

// convert entity to model
func (m *TransactionModel) FromEntity(e *entity.Transaction) *TransactionModel {
	return &TransactionModel{
		ID: e.ID,
		PublicId: e.PublicId,		
		KonsumenTenorLimitPublicId: e.KonsumenTenorLimitPublicId,
		ProductPublicId: e.ProductPublicId,
		NomorKontrak: e.NomorKontrak,
		HargaOtr: e.HargaOtr,
		AdminFee: e.AdminFee,
		JumlahCicilan: e.JumlahCicilan,
		JumlahBunga: e.JumlahBunga,
		NamaAset: e.NamaAset,
		CreatedAt: e.CreatedAt,
		CreatedBy: e.CreatedBy,
		UpdatedAt: e.UpdatedAt,
		UpdatedBy: e.UpdatedBy,
	}
}

// reverse model to entity
func (m *TransactionModel) ToEntity() *entity.Transaction {
	return &entity.Transaction{
		ID: m.ID,
		PublicId: m.PublicId,		
		KonsumenTenorLimitPublicId: m.KonsumenTenorLimitPublicId,
		ProductPublicId: m.ProductPublicId,
		NomorKontrak: m.NomorKontrak,
		HargaOtr: m.HargaOtr,
		AdminFee: m.AdminFee,
		JumlahCicilan: m.JumlahCicilan,
		JumlahBunga: m.JumlahBunga,
		NamaAset: m.NamaAset,
		CreatedAt: m.CreatedAt,
		CreatedBy: m.CreatedBy,
		UpdatedAt: m.UpdatedAt,
		UpdatedBy: m.UpdatedBy,
	}
}