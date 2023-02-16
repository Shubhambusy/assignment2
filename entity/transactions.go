package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id              uuid.UUID
	Time            time.Time
	Date            time.Time
	TransactionType string  // deposits, withdraw, transfer  -------   0, 1, 2
	SenderAccount   int64
	RecipientAcount int64
	Amount          int64
}
