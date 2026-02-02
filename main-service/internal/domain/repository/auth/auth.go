package auth

import (
	"context"
	entity "ptxyz/main-service/internal/domain/entity/auth"
)

type AuthRepository interface {
	VerifyCredential(ctx context.Context, nik string, password string) (*entity.AuthIdentity, error)
}
