package auth

// DTO
type VerifyInput struct {
	Nik    string
	Password string
}

type VerifyOutput struct {
	UserPublicId string
}