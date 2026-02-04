package konsumentenorlimit

import (
	"context"
	entity "ptxyz/transaction-service/internal/domain/entity/konsumen_tenor_limit"
)

type KonsumenTenorLimitRepository interface {
	UpdateBalance(ctx context.Context, e *entity.KTLUpdateBalance) error
}