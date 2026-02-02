package user

import (
	"mime/multipart"
)

type CreateUserRequest struct {
	Nik           string                `form:"nik" binding:"required,min=8,max=20"`
	Password	  string				`form:"password" binding:"required,min=6"`
	FullName      string                `form:"fullName" binding:"required"`
	LegalName     string                `form:"legalName" binding:"required"`
	TempatLahir   string                `form:"tempatLahir" binding:"required"`
	TanggalLahir  string                `form:"tanggalLahir" binding:"required"` // YYYY-MM-DD
	Gaji          string                `form:"gaji" binding:"required"`
	FotoKtp       *multipart.FileHeader `form:"fotoKtp" binding:"required"`
	FotoSelfie    *multipart.FileHeader `form:"fotoSelfie" binding:"required"`
}