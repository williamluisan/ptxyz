package user

import (
	"mime/multipart"
	"time"

	"github.com/shopspring/decimal"
)

type CreateUserRequest struct {
	Nik				string
	FullName		string
	LegalName		string
	TempatLahir		string
	TanggalLahir	time.Time
	Gaji			decimal.Decimal
	FotoKtp			*multipart.FileHeader
	FotoSelfie		*multipart.FileHeader
}