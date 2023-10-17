package models

import "github.com/google/uuid"

const (
	AccountTypeSalary  = "Salary"
	AccountTypeSaving  = "Saving"
	AccountTypeCurrent = "Current"
)

// Account model
type Account struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	AccountType string    `json:"account_type"`
	Balance     float64   `json:"balance"`
}
