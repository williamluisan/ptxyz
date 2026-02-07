package konsumentenorlimit

import (
	"context"
	entity "ptxyz/transaction-service/internal/domain/entity/konsumentenorlimit"
)

type KonsumenTenorLimitRepository interface {
	UpdateBalance(ctx context.Context, e *entity.KTLUpdateBalance) error
}