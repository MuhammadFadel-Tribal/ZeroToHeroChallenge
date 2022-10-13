package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type TransactionModel struct {
	bun.BaseModel `bun:"table:transaction"`
	ID            uuid.UUID `bun:"id,notnull,pk,type:uuid,default:gen_random_uuid()"`
	Amount        float64   `bun:"amount,notnull"`
	Currency      string    `bun:"currency,notnull"`
	CreatedAt     time.Time `bun:"createdat,notnull"`
}

type CreateTransactionDto struct {
	Amount   float64 `json:"Amount" validate:"required"`
	Currency string  `json:"Currency" validate:"required"`
}

type TransactionDetailsDto struct {
	ID        uuid.UUID `json:"Id"`
	Amount    float64   `json:"Amount"`
	Currency  string    `json:"Currency"`
	CreatedAt time.Time `json:"createdAt"`
}
