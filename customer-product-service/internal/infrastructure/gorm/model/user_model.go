package model

import (
	"time"

	entity "ptxyz/customer-product-service/internal/domain/entity/user"

	"github.com/shopspring/decimal"
)

type UserModel struct {
	ID				uint64			`gorm:"primaryKey;unique;not null;"`
	PublicId		string			`gorm:"type:varchar(255);not null;"`
	Nik				string			`gorm:"type:varchar(100);not null;"`
	Password		string			`gorm:"type:varchar(255);not null;"`
	FullName		string			`gorm:"type:varchar(255);not null;"`
	LegalName		string			`gorm:"type:varchar(255);not null;"`
	TempatLahir		string			`gorm:"type:varchar(50);not null;"`
	TanggalLahir	time.Time		`gorm:"type:date;not null;"`
	Gaji			decimal.Decimal	`gorm:"type:decimal(15,2);not null;"`	
	FotoKtp			string			`gorm:"type:varchar(255);not null;"`
	FotoSelfie		string			`gorm:"type:varchar(255);not null;"`
	CreatedAt		time.Time
	CreatedBy		int
	UpdatedAt		time.Time
	UpdatedBy		int
}

func (UserModel) TableName() string {
	return "users"
}

// convert entity to model
func (m *UserModel) FromEntity(u *entity.User) *UserModel {
	return &UserModel{
		ID: u.ID,
		PublicId: u.PublicId,
		Nik: u.Nik,
		Password: u.Password,				
		FullName: u.FullName,		
		LegalName: u.LegalName,		
		TempatLahir: u.TempatLahir,		
		TanggalLahir: u.TanggalLahir,	
		Gaji: u.Gaji,			
		FotoKtp: u.FotoKtp,		
		FotoSelfie: u.FotoSelfie,
		CreatedBy: 1, // get from logged in user
	}
}

// reverse model to entity
func (m *UserModel) ToEntity() *entity.User {
	return &entity.User{
		// ID: m.ID,
		PublicId: m.PublicId,
		Nik: m.Nik,		
		Password: m.Password,				
		FullName: m.FullName,		
		LegalName: m.LegalName,		
		TempatLahir: m.TempatLahir,		
		TanggalLahir: m.TanggalLahir,	
		Gaji: m.Gaji,			
		FotoKtp: m.FotoKtp,		
		FotoSelfie: m.FotoSelfie,
		CreatedAt: m.CreatedAt,
	}
}