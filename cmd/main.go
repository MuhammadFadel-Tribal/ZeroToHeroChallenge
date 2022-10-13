package main

import (
	"zerotoherochallenge/config"
	"zerotoherochallenge/internal/adaptors/api"
	"zerotoherochallenge/internal/adaptors/db"
	"zerotoherochallenge/internal/adaptors/stream"
	"zerotoherochallenge/internal/repositories"
	"zerotoherochallenge/internal/services/Transaction"

	"github.com/go-playground/validator/v10"
)

func main() {

	// Logger
	logger := config.NewLogger()
	defer config.CloseLogger(logger)

	// Configs
	configs := config.LoadConfig(logger)

	// Database Connection
	database := db.ConnectDatabase(logger, configs.Database)

	httpServer := api.NewHTTPServer(logger, configs.Server)

	// Validator
	validate := validator.New()

	transactionRepo := repositories.DatabaseInitializer(logger, database)

	transactionSvc := Transaction.TransactionServiceInitializer(logger, transactionRepo)

	api.TransactionControllerInitializer(httpServer, logger, validate, transactionSvc)

	// -- End dependency injection section --

	//kafka
	//stream.InitializeKafka(configs)
	go stream.NewKafkaConsumer(logger)

	// Let the party started!
	httpServer.Start()

}
