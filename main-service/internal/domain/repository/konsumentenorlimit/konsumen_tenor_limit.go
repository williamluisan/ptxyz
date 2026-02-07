package konsumentenorlimit

import (
	"context"
	entity "ptxyz/main-service/internal/domain/entity/konsumentenorlimit"
)

type KonsumenTenorLimitRepository interface {
	GetByPublicId(ctx context.Context, publicId string) (*entity.KonsumenTenorLimit, error)
}