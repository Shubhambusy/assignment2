package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id              uuid.UUID `pg:"type:uuid"`
	Timestamp       time.Time
	Type            string // cash_deposit, cash_withdraw, account_transfer
	SenderAccount   int64  // for cash transaction account number = 0
	RecipientAcount int64  // for cash transaction account number = 0
	Amount          float64
}

// transaction type can be optimized by using enum,
// but it was not properly supported by go-pg

// Also, uuid can stores timestamp by default
// se we can fetch it from there itself
