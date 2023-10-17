package controllers

import (
	"database/sql"
	"log"
	"vbank/internal/models"

	"github.com/google/uuid"
)

type AddressController struct {
	db *sql.DB
}

func NewAddressController() *AddressController {
	return &AddressController{
		db: DB,
	}
}

func (ac *AddressController) CreateAddress(a *models.Address) uuid.UUID {
	err := ac.db.QueryRow(
		`INSERT INTO addresses (id, user_id, street_address, city, state, postal_code, address_type)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`, a.ID, a.UserID, a.StreetAddress, a.City, a.State, a.PostalCode, a.AddressType).Scan(&a.ID)
	if err != nil {
		log.Fatal(err)
	}
	return a.ID
}

func (ac *AddressController) UpdateAddress(a *models.Address) {
	_, err := ac.db.Exec(
		`UPDATE addresses
		SET user_id = $2, street_address = $3, city = $4, state = $5, postal_code = $6, address_type = $7
		WHERE id = $1`, a.ID, a.UserID, a.StreetAddress, a.City, a.State, a.PostalCode, a.AddressType)
	if err != nil {
		log.Fatal(err)
	}
}

func (ac *AddressController) GetAddressByID(id uuid.UUID) models.Address {
	var address models.Address
	err := ac.db.QueryRow(
		`SELECT id, user_id, street_address, city, state, postal_code, address_type
		FROM addresses WHERE id = $1`, id).Scan(&address.ID, &address.UserID, &address.StreetAddress, &address.City, &address.State, &address.PostalCode, &address.AddressType)
	if err != nil {
		log.Fatal(err)
	}
	return address
}

func (ac *AddressController) GetAddressesByUserID(userID uuid.UUID) []models.Address {
	var addresses []models.Address

	rows, err := ac.db.Query("SELECT id, user_id, street_address, city, state, postal_code, address_type FROM addresses WHERE user_id = $1", userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var address models.Address
		err := rows.Scan(&address.ID, &address.UserID, &address.StreetAddress, &address.City, &address.State, &address.PostalCode, &address.AddressType)
		if err != nil {
			log.Fatal(err)
		}
		addresses = append(addresses, address)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return addresses
}
func (ac *AddressController) DeleteAddressByUserID(userID uuid.UUID) {
	_, err := ac.db.Exec("DELETE FROM addresses WHERE user_id = $1", userID)
	if err != nil {
		log.Fatal(err)
	}
}

func (ac *AddressController) DeleteAddressByAddressID(addressID uuid.UUID) {
	_, err := ac.db.Exec("DELETE FROM addresses WHERE id = $1", addressID)
	if err != nil {
		log.Fatal(err)
	}
}

func (ac *AddressController) UpdateAddressByUserID(userID uuid.UUID, updatedAddress models.Address) {
	_, err := ac.db.Exec(
		"UPDATE addresses "+
			"SET street_address = $2, city = $3, state = $4, postal_code = $5, address_type = $6 "+
			"WHERE user_id = $1", userID, updatedAddress.StreetAddress, updatedAddress.City, updatedAddress.State, updatedAddress.PostalCode, updatedAddress.AddressType)
	if err != nil {
		log.Fatal(err)
	}
}
