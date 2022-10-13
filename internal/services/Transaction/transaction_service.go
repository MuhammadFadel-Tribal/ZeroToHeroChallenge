package Transaction

import (
	"context"
	"github.com/google/uuid"
	"time"
	"zerotoherochallenge/internal/adaptors/stream"
	"zerotoherochallenge/internal/models"
	"zerotoherochallenge/internal/repositories"

	"github.com/dranikpg/dto-mapper"
	"go.uber.org/zap"
)

type TransactionService struct {
	log *zap.SugaredLogger
	db  repositories.ITransactionRepository
}

func TransactionServiceInitializer(logger *zap.SugaredLogger, transactionRepo repositories.ITransactionRepository) *TransactionService {
	return &TransactionService{
		log: logger,
		db:  transactionRepo,
	}
}

func (ds *TransactionService) FindAllTransactions(ctx context.Context) (*[]models.TransactionDetailsDto, bool) {

	err, Result := ds.db.FindAll(ctx)
	if err != nil {
		ds.log.Errorf("Error while retrieving transactions: %s", err)
		return nil, false
	}

	var data = &[]models.TransactionDetailsDto{}

	dto.Map(data, Result)
	ds.log.Infof("Retrieving all transactions")

	return data, true
}

func (ds *TransactionService) FindTransactionDetails(ctx context.Context, id string) (*models.TransactionDetailsDto, bool) {

	result, ok := ds.db.FindById(ctx, id)
	if !ok {
		ds.log.Errorf("This transaction is not exist")
		return nil, false
	}
	ds.log.Infof("Finding details for transaction --> %s", id)
	response := &models.TransactionDetailsDto{}
	dto.Map(response, result)

	return response, true
}

func (ds *TransactionService) AddTransaction(ctx context.Context, request *models.CreateTransactionDto) (*models.TransactionDetailsDto, bool) {
	var trans = &models.TransactionModel{}
	dto.Map(trans, request)
	trans.ID = uuid.New()
	trans.CreatedAt = time.Now().UTC()

	err, result := ds.db.Create(ctx, trans)

	if err != nil {
		ds.log.Errorf("Error with creating transaction: %s", err)
		return nil, false
	}

	var transResponse = &models.TransactionDetailsDto{}

	dto.Map(trans, result)
	dto.Map(transResponse, trans)
	ds.log.Infof("Transaction created!")

	stream.NewKafkaProducer(trans)
	return transResponse, true
}
