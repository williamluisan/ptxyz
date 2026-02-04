package transaction

import (
	"context"
	entity "ptxyz/transaction-service/internal/domain/entity/transaction"
	repository "ptxyz/transaction-service/internal/domain/repository/transaction"
)

type TransactionService interface {
	Create(ctx context.Context, input *entity.TransactionInput) error
}

type transactionServiceImpl	struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) TransactionService {
	if repo == nil {
		panic("transaction repository cannot be nil")
	}
	return &transactionServiceImpl{
		repo: repo,
	}
}

func (i *transactionServiceImpl) Create(ctx context.Context, input *entity.TransactionInput) error {
	data := input

	// check if konsumen tenor limit exists
	// ...

	// check if product exists
	// ...

	// insert
	newTransaction := entity.NewTransaction(data)
	if err := i.repo.Create(ctx, newTransaction); err != nil {
		return err
	}

	// update tenor balance (rollback if fail)
	// ..

	return nil
}