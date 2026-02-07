package transaction

import "time"

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
