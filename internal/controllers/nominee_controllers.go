package controllers

import (
	"database/sql"
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

func (nc *NomineeController) CreateNominee(n *models.Nominee) (uuid.UUID, error) {
	err := nc.db.QueryRow(
		`INSERT INTO nominees (id, user_id, name, relationship)
		VALUES ($1, $2, $3, $4)
		RETURNING id`, n.ID, n.UserID, n.Name, n.Relationship).Scan(&n.ID)
	if err != nil {
		return uuid.Nil, err
	}
	return n.ID, nil
}

func (nc *NomineeController) UpdateNominee(n *models.Nominee) error {
	_, err := nc.db.Exec(
		`UPDATE nominees
		SET user_id = $2, name = $3, relationship = $4
		WHERE id = $1`, n.ID, n.UserID, n.Name, n.Relationship)
	if err != nil {
		return err
	}
	return nil
}

func (nc *NomineeController) GetNomineeByID(id uuid.UUID) (models.Nominee, error) {
	var nominee models.Nominee
	err := nc.db.QueryRow(
		`SELECT id, user_id, name, relationship
		FROM nominees WHERE id = $1`, id).Scan(&nominee.ID, &nominee.UserID, &nominee.Name, &nominee.Relationship)
	if err != nil {
		return nominee, err
	}
	return nominee, nil
}

func (nc *NomineeController) GetNomineesByUserID(userID uuid.UUID) ([]models.Nominee, error) {
	var nominees []models.Nominee

	rows, err := nc.db.Query("SELECT id, user_id, name, relationship FROM nominees WHERE user_id = $1", userID)
	if err != nil {
		return nominees, err
	}
	defer rows.Close()

	for rows.Next() {
		var nominee models.Nominee
		err := rows.Scan(&nominee.ID, &nominee.UserID, &nominee.Name, &nominee.Relationship)
		if err != nil {
			return nominees, err
		}
		nominees = append(nominees, nominee)
	}

	if err := rows.Err(); err != nil {
		return nominees, err
	}

	return nominees, nil
}
