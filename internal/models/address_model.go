package models

import "github.com/google/uuid"

const (
	AddressTypeRent      = "Rent"
	AddressTypeCurrent   = "Current"
	AddressTypePermanent = "Permanent"
)

// Address model
type Address struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	StreetAddress string    `json:"street_address"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	PostalCode    string    `json:"postal_code"`
	AddressType   string    `json:"address_type"`
}
