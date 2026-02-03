package user

import (
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/shopspring/decimal"
)

type User struct {
	ID				uint64
	PublicId		string
	Nik				string
	Password		string
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
	Password		string
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
		PublicId: ulid.Make().String(),	
		Password: input.Password,		
		FullName: input.FullName,		
		LegalName: input.LegalName,		
		TempatLahir: input.TempatLahir,		
		TanggalLahir: input.TanggalLahir,	
		Gaji: input.Gaji,			
		FotoKtp: input.FotoKtp,		
		FotoSelfie: input.FotoSelfie,
		CreatedAt: time.Now(),
		CreatedBy: 1,		
	}
}