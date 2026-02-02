package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	secret []byte
	ttl    time.Duration
}

func NewTokenService(secret string, ttl time.Duration) *TokenService {
	return &TokenService{
		secret: []byte(secret),
		ttl: ttl,
	}
}

func (t *TokenService) Generate(userPublicId string) (string, error) {

	claims := jwt.MapClaims{
		"sub": userPublicId,
		"exp": time.Now().Add(t.ttl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(t.secret)
}
