package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	TxnTypeCredit = "Credit"
	TxnTypeDebit  = "Debit"
)

// Transaction model
type Transaction struct {
	ID              uuid.UUID `json:"id"`
	AccountID       uuid.UUID `json:"account_id"`
	TransactionType string    `json:"transaction_type"`
	Amount          float64   `json:"amount"`
	TransactionDate time.Time `json:"transaction_date"`
}
