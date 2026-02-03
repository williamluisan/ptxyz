package entity

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Transaction struct {
	ID							uint64
	PublicId					string
	KonsumenTenorLimitPublicId	string
	ProductPublicId				string
	NomorKontrak				string
	HargaOtr					string
	AdminFee					string
	JumlahCicilan				string
	JumlahBunga					string
	NamaAset					string
	CreatedAt					time.Time
	CreatedBy					int
	UpdatedAt					time.Time
	UpdatedBy					int
}

// DTO
type TransactionInput struct {
	KonsumenTenorLimitPublicId	string
	ProductPublicId				string
	NomorKontrak				string
	HargaOtr					string
	AdminFee					string
	JumlahCicilan				string
	JumlahBunga					string
	NamaAset					string
}

func NewTransaction(input *TransactionInput) *Transaction {
	return &Transaction{
		PublicId: ulid.Make().String(),
		KonsumenTenorLimitPublicId: input.KonsumenTenorLimitPublicId,	
		ProductPublicId: input.ProductPublicId,				
		NomorKontrak: input.NomorKontrak,				
		HargaOtr: input.HargaOtr,					
		AdminFee: input.AdminFee,					
		JumlahCicilan: input.JumlahCicilan,				
		JumlahBunga: input.JumlahBunga,					
		NamaAset: input.NamaAset,					
		CreatedAt: time.Now(),
		CreatedBy: 1,
	}
}