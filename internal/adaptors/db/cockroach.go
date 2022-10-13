package db

import (
	"zerotoherochallenge/config"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
	"go.uber.org/zap"
)

func ConnectDatabase(log *zap.SugaredLogger, config config.DatabaseConfigurations) *bun.DB {

	parsedCfg, err := pgx.ParseConfig(config.Dsn)

	if err != nil {
		log.Fatalf("DB connection error -> %v", err)
	}

	parsedCfg.PreferSimpleProtocol = true

	sqlDb := stdlib.OpenDB(*parsedCfg)

	sqlDb.SetMaxOpenConns(config.Pool)

	err = sqlDb.Ping()
	if err != nil {
		log.Fatalf("DB connection error -> %v", err)
	}

	db := bun.NewDB(sqlDb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook())

	log.Infof("Database connected successfully. Connections opened: %d", db.Stats().OpenConnections)

	return db
}
