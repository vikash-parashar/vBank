package controllers

import (
	"database/sql"
	"log"
	"vbank/internal/models"

	"github.com/google/uuid"
)

type AccountController struct {
	db *sql.DB
}

func NewAccountController() *AccountController {
	return &AccountController{
		db: DB,
	}
}

func (ac *AccountController) CreateAccount(a *models.Account) uuid.UUID {
	err := ac.db.QueryRow(
		`INSERT INTO accounts (id, user_id, account_type, balance)
		VALUES ($1, $2, $3, $4)
		RETURNING id`, a.ID, a.UserID, a.AccountType, a.Balance).Scan(&a.ID)
	if err != nil {
		log.Fatal(err)
	}
	return a.ID
}

func (ac *AccountController) UpdateAccount(a *models.Account) {
	_, err := ac.db.Exec(
		`UPDATE accounts
		SET user_id = $2, account_type = $3, balance = $4
		WHERE id = $1`, a.ID, a.UserID, a.AccountType, a.Balance)
	if err != nil {
		log.Fatal(err)
	}
}

func (ac *AccountController) GetAccountByID(id uuid.UUID) models.Account {
	var account models.Account
	err := ac.db.QueryRow(
		`SELECT id, user_id, account_type, balance
		FROM accounts WHERE id = $1`, id).Scan(&account.ID, &account.UserID, &account.AccountType, &account.Balance)
	if err != nil {
		log.Fatal(err)
	}
	return account
}

func (ac *AccountController) GetAccountsByUserID(userID uuid.UUID) []models.Account {
	var accounts []models.Account

	rows, err := ac.db.Query("SELECT id, user_id, account_type, balance FROM accounts WHERE user_id = $1", userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var account models.Account
		err := rows.Scan(&account.ID, &account.UserID, &account.AccountType, &account.Balance)
		if err != nil {
			log.Fatal(err)
		}
		accounts = append(accounts, account)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return accounts
}
