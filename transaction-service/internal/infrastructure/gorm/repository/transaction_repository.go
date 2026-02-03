package repository

import (
	entity "ptxyz/transaction-service/internal/domain/entity/transaction"
	repository "ptxyz/transaction-service/internal/domain/repository/transaction"
	gormModel "ptxyz/transaction-service/internal/infrastructure/gorm/model"

	"context"

	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) repository.TransactionRepository {
	return &transactionRepository{db: db}
}

var mTransactionModel gormModel.TransactionModel

func (r *transactionRepository) Create(ctx context.Context, e *entity.Transaction) error {
	data := mTransactionModel.FromEntity(e)

	if err := gorm.G[gormModel.TransactionModel](r.db).Create(ctx, data); err != nil {
		return err
	}

	return nil
}