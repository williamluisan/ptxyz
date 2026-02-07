package transaction

import (
	"context"

	entity "ptxyz/main-service/internal/domain/entity/transaction"
)

type TransactionRepository interface {
	Create(ctx context.Context, e *entity.Transaction) error
}