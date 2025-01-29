package repository

import (
	"context"
	"ladipage_server/core/adapters"
	"ladipage_server/core/domain"

	"gorm.io/gorm"
)

type TransactionHelperRepository struct {
	db *adapters.Pgsql
}

func NewRepositoryTransaction(db *adapters.Pgsql) domain.RepositoryTransactionHelper {
	return &TransactionHelperRepository{
		db: db,
	}
}

func (t *TransactionHelperRepository) Transaction(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) error) error {
	return t.db.DB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(ctx, tx)
	})
}
