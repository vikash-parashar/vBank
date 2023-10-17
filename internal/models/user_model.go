package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	UserTypeCustomer = "Customer"
	UserTypeAdmin    = "Admin"
)
const (
	GenderMale   = "Male"
	GenderFemale = "Female"
)

// User model
type User struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Gender      string    `json:"gender"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
}
