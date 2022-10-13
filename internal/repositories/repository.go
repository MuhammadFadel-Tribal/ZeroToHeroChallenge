package repositories

import (
	"context"
	"zerotoherochallenge/internal/models"
)

type ITransactionRepository interface {
	FindAll(ctx context.Context) (error, []models.TransactionModel)
	FindById(ctx context.Context, transactionID string) (models.TransactionModel, bool)
	Create(ctx context.Context, transaction *models.TransactionModel) (error, *models.TransactionDetailsDto)
}
