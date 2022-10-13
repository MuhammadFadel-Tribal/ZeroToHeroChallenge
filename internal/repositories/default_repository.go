package repositories

import (
	"context"
	"zerotoherochallenge/internal/models"

	"github.com/dranikpg/dto-mapper"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

type DatabaseContext struct {
	log *zap.SugaredLogger
	db  *bun.DB
}

// DatabaseInitializer Creates a new repository
func DatabaseInitializer(logger *zap.SugaredLogger, conn *bun.DB) *DatabaseContext {
	return &DatabaseContext{
		log: logger,
		db:  conn,
	}
}

func (d *DatabaseContext) FindAll(ctx context.Context) (error, []models.TransactionModel) {
	var model = &[]models.TransactionModel{}

	err := d.db.NewSelect().
		Model(model).
		Scan(ctx)

	if err != nil {
		d.log.Errorf("Error while retrieving transactions: %s", err)
		return err, nil
	}

	d.log.Infof("Retrieving all transactions")
	return nil, *model
}

func (d *DatabaseContext) FindById(ctx context.Context, transactionID string) (models.TransactionModel, bool) {
	var model = &models.TransactionModel{}

	err := d.db.NewSelect().
		Model(model).
		Where("id = ?", transactionID).
		Scan(ctx)

	if err != nil {
		return *model, false
	}

	return *model, true

}

func (d *DatabaseContext) Create(ctx context.Context, transaction *models.TransactionModel) (error, *models.TransactionDetailsDto) {
	_, err := d.db.NewInsert().Model(transaction).Exec(ctx)
	if err != nil {
		return err, nil
	}

	var transactionModel = &models.TransactionDetailsDto{}
	dto.Map(transactionModel, transaction)

	return nil, transactionModel
}
