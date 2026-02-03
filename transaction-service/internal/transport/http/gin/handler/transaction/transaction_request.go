package transaction

type CreateTransactionRequest struct {
	KonsumenTenorLimitPublicId	string	`json:"konsumen_tenor_limit_public_id" binding:"required"`
	ProductPublicId				string	`json:"product_public_id" binding:"required"`
	NomorKontrak				string	`json:"nomor_kontrak" binding:"required,alphanum"`
	HargaOtr					string	`json:"harga_otr" binding:"required,numeric"`
	AdminFee					string	`json:"admin_fee" binding:"required,numeric"`
	JumlahCicilan				string	`json:"jumlah_cicilan" binding:"required,numeric"`
	JumlahBunga					string	`json:"jumlah_bunga" binding:"required,numeric"`
	NamaAset					string  `json:"nama_aset" binding:"required"`
}