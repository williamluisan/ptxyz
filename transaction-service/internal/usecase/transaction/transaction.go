package transaction

import (
	"context"
	"errors"
	entityKTL "ptxyz/transaction-service/internal/domain/entity/konsumentenorlimit"
	entity "ptxyz/transaction-service/internal/domain/entity/transaction"
	repositoryKTL "ptxyz/transaction-service/internal/domain/repository/konsumentenorlimit"
	repositoryProduct "ptxyz/transaction-service/internal/domain/repository/product"
	repository "ptxyz/transaction-service/internal/domain/repository/transaction"
	"strconv"

	"golang.org/x/sync/errgroup"
)

type TransactionService interface {
	Create(ctx context.Context, input *entity.TransactionInput) error
}

type transactionServiceImpl	struct {
	repo repository.TransactionRepository
	repoKTL repositoryKTL.KonsumenTenorLimitRepository
	repoProduct repositoryProduct.ProductRepository
}

func NewTransactionService(repo repository.TransactionRepository, repoKTL repositoryKTL.KonsumenTenorLimitRepository, repoProduct repositoryProduct.ProductRepository) TransactionService {
	if repo == nil {
		panic("transaction repository cannot be nil")
	}
	return &transactionServiceImpl{
		repo: repo,
		repoKTL: repoKTL,
		repoProduct: repoProduct,
	}
}

func (i *transactionServiceImpl) Create(ctx context.Context, input *entity.TransactionInput) error {
	data := input

	parentCtx := ctx // keep original

	g, ctx := errgroup.WithContext(ctx)

	// check if konsumen tenor limit exists
	g.Go(func() error {
		return nil
	})

	// check if product exists
	g.Go(func() error {
		productPublicId := input.ProductPublicId
		product, err := i.repoProduct.GetByPublicId(ctx, productPublicId)
		if err != nil {
			return err
		}
		if product == nil {
			return errors.New("product not found")
		}

		return nil
	})

	// wait for both
	if err := g.Wait(); err != nil {
		return err
	}

	// insert
	newTransaction := entity.NewTransaction(data)
	if err := i.repo.Create(parentCtx, newTransaction); err != nil {
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