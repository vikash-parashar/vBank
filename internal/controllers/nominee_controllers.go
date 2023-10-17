package controllers

import (
	"database/sql"
	"log"
	"vbank/internal/models"

	"github.com/google/uuid"
)

type NomineeController struct {
	db *sql.DB
}

func NewNomineeController() *NomineeController {
	return &NomineeController{
		db: DB,
	}
}

func (nc *NomineeController) CreateNominee(n *models.Nominee) uuid.UUID {
	err := nc.db.QueryRow(
		`INSERT INTO nominees (id, user_id, name, relationship)
		VALUES ($1, $2, $3, $4)
		RETURNING id`, n.ID, n.UserID, n.Name, n.Relationship).Scan(&n.ID)
	if err != nil {
		log.Fatal(err)
	}
	return n.ID
}

func (nc *NomineeController) UpdateNominee(n *models.Nominee) {
	_, err := nc.db.Exec(
		`UPDATE nominees
		SET user_id = $2, name = $3, relationship = $4
		WHERE id = $1`, n.ID, n.UserID, n.Name, n.Relationship)
	if err != nil {
		log.Fatal(err)
	}
}

func (nc *NomineeController) GetNomineeByID(id uuid.UUID) models.Nominee {
	var nominee models.Nominee
	err := nc.db.QueryRow(
		`SELECT id, user_id, name, relationship
		FROM nominees WHERE id = $1`, id).Scan(&nominee.ID, &nominee.UserID, &nominee.Name, &nominee.Relationship)
	if err != nil {
		log.Fatal(err)
	}
	return nominee
}

func (nc *NomineeController) GetNomineesByUserID(userID uuid.UUID) []models.Nominee {
	var nominees []models.Nominee

	rows, err := nc.db.Query("SELECT id, user_id, name, relationship FROM nominees WHERE user_id = $1", userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var nominee models.Nominee
		err := rows.Scan(&nominee.ID, &nominee.UserID, &nominee.Name, &nominee.Relationship)
		if err != nil {
			log.Fatal(err)
		}
		nominees = append(nominees, nominee)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return nominees
}
