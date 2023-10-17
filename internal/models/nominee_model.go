package models

import "github.com/google/uuid"

const (
	RelationWithNomineeFather   = "Father"
	RelationWithNomineeMother   = "Mother"
	RelationWithNomineeBrother  = "Brother"
	RelationWithNomineeSister   = "Sister"
	RelationWithNomineeSon      = "Son"
	RelationWithNomineeDaughter = "Daughter"
)

// Nominee model
type Nominee struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	Name         string    `json:"name"`
	Relationship string    `json:"relationship"`
}
