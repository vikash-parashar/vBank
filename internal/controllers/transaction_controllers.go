package controllers

import (
	"database/sql"
	"log"
	"vbank/internal/models"

	"github.com/google/uuid"
)

type TransactionController struct {
	db *sql.DB
}

func NewTransactionController() *TransactionController {
	return &TransactionController{
		db: DB,
	}
}

func (tc *TransactionController) CreateTransaction(t *models.Transaction) uuid.UUID {
	err := tc.db.QueryRow(
		`INSERT INTO transactions (id, account_id, transaction_type, amount, transaction_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`, t.ID, t.AccountID, t.TransactionType, t.Amount, t.TransactionDate).Scan(&t.ID)
	if err != nil {
		log.Fatal(err)
	}
	return t.ID
}

func (tc *TransactionController) UpdateTransaction(t *models.Transaction) {
	_, err := tc.db.Exec(
		`UPDATE transactions
		SET account_id = $2, transaction_type = $3, amount = $4, transaction_date = $5
		WHERE id = $1`, t.ID, t.AccountID, t.TransactionType, t.Amount, t.TransactionDate)
	if err != nil {
		log.Fatal(err)
	}
}

func (tc *TransactionController) GetTransactionByID(id uuid.UUID) models.Transaction {
	var transaction models.Transaction
	err := tc.db.QueryRow(
		`SELECT id, account_id, transaction_type, amount, transaction_date
		FROM transactions WHERE id = $1`, id).Scan(&transaction.ID, &transaction.AccountID, &transaction.TransactionType, &transaction.Amount, &transaction.TransactionDate)
	if err != nil {
		log.Fatal(err)
	}
	return transaction
}

func (tc *TransactionController) GetTransactionsByAccount(accountID uuid.UUID) []models.Transaction {
	var transactions []models.Transaction

	rows, err := tc.db.Query("SELECT id, account_id, transaction_type, amount, transaction_date FROM transactions WHERE account_id = $1", accountID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(&transaction.ID, &transaction.AccountID, &transaction.TransactionType, &transaction.Amount, &transaction.TransactionDate)
		if err != nil {
			log.Fatal(err)
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return transactions
}

func (tc *TransactionController) GetAllTransactionsByUserID(userID uuid.UUID) []models.Transaction {
	var transactions []models.Transaction

	rows, err := tc.db.Query("SELECT t.id, t.account_id, t.transaction_type, t.amount, t.transaction_date "+
		"FROM transactions t "+
		"INNER JOIN accounts a ON t.account_id = a.id "+
		"WHERE a.user_id = $1", userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(&transaction.ID, &transaction.AccountID, &transaction.TransactionType, &transaction.Amount, &transaction.TransactionDate)
		if err != nil {
			log.Fatal(err)
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return transactions
}
