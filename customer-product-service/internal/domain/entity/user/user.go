package user

import (
	"time"

	"github.com/shopspring/decimal"
)

type User struct {
	ID				uint64
	PublicId		string
	Nik				string
	FullName		string
	LegalName		string
	TempatLahir		string
	TanggalLahir	time.Time
	Gaji			decimal.Decimal
	FotoKtp			string
	FotoSelfie		string
	CreatedAt		time.Time
	CreatedBy		int
	UpdatedAt		time.Time
	UpdatedBy		int
}

// DTO (Data Transfer Object)
type UserInput struct {
	Nik				string
	FullName		string
	LegalName		string
	TempatLahir		string
	TanggalLahir	time.Time
	Gaji			decimal.Decimal
	FotoKtp			string
	FotoSelfie		string
}

func NewUser(input *UserInput) *User {
	return &User{
		Nik: input.Nik,				
		FullName: input.FullName,		
		LegalName: input.LegalName,		
		TempatLahir: input.TempatLahir,		
		TanggalLahir: input.TanggalLahir,	
		Gaji: input.Gaji,			
		FotoKtp: input.FotoKtp,		
		FotoSelfie: input.FotoSelfie,		
	}
}