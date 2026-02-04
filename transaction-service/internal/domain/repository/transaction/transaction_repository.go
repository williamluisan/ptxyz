package repository

import (
	"context"
	entity "ptxyz/transaction-service/internal/domain/entity/transaction"
)

type TransactionRepository interface {
	Create(ctx context.Context, e *entity.Transaction) error
}

type TransactionRepositoryServiceRepository interface {
	UpdateBalance(ctx context.Context, e *entity.TransactionInput) error
}