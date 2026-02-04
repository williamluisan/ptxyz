package transaction

import (
	"context"
	entityKTL "ptxyz/transaction-service/internal/domain/entity/konsumen_tenor_limit"
	entity "ptxyz/transaction-service/internal/domain/entity/transaction"
	repositoryKTL "ptxyz/transaction-service/internal/domain/repository/konsumen_tenor_limit"
	repository "ptxyz/transaction-service/internal/domain/repository/transaction"
	"strconv"
	"sync"
)

type TransactionService interface {
	Create(ctx context.Context, input *entity.TransactionInput) error
}

type transactionServiceImpl	struct {
	mu *sync.Mutex
	repo repository.TransactionRepository
	repoKTL repositoryKTL.KonsumenTenorLimitRepository
}

func NewTransactionService(repo repository.TransactionRepository, repoKTL repositoryKTL.KonsumenTenorLimitRepository) TransactionService {
	if repo == nil {
		panic("transaction repository cannot be nil")
	}
	return &transactionServiceImpl{
		repo: repo,
		repoKTL: repoKTL,
		mu: &sync.Mutex{},
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
	limit := float64(1500000)
	usedAmount, _ := strconv.ParseFloat(input.HargaOtr, 64)
	adminFee, _ := strconv.ParseFloat(input.AdminFee, 64)
	totalUsedAmount := usedAmount + adminFee
	balance, _ := i.calculateUsedAmount(limit, totalUsedAmount)
	KTLUpdateBalance := &entityKTL.KTLUpdateBalance{
		PublicId: input.KonsumenTenorLimitPublicId,
		UsedAmount: usedAmount,
		Balance: balance,
	}
	i.repoKTL.UpdateBalance(ctx, KTLUpdateBalance);

	return nil
}

func (i *transactionServiceImpl) calculateUsedAmount(limit float64, usedAmount float64) (float64, error) {
	balance := limit - usedAmount

	return balance, nil
}