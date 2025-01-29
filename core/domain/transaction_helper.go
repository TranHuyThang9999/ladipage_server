package domain

import (
	"context"

	"gorm.io/gorm"
)

type RepositoryTransactionHelper interface {
	Transaction(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) error) error
}
