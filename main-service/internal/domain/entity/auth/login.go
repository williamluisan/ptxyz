package auth

// DTO
type LoginInput struct {
	Nik    string
	Password string
}

type LoginOutput struct {
	AccessToken string
}