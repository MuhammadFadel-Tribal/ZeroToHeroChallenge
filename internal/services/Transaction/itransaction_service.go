package Transaction

import (
	"context"
	"zerotoherochallenge/internal/models"
)

type ITransactionService interface {
	FindAllTransactions(ctx context.Context) ([]*models.TransactionDetailsDto, bool)
	FindTransactionDetails(ctx context.Context, id string) (*models.TransactionDetailsDto, bool)
	AddTransaction(ctx context.Context, request *models.CreateTransactionDto) (*models.TransactionDetailsDto, bool)
}
