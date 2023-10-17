package controllers

import (
	"database/sql"
	"log"
	"vbank/internal/db"
	"vbank/internal/models"

	"github.com/google/uuid"
)

var DB = db.GetDB()

type UserController struct {
	db *sql.DB
}

func NewUserController() *UserController {
	return &UserController{
		db: DB,
	}
}

func (uc *UserController) CreateUserTable() {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY,
		first_name VARCHAR(255),
		last_name VARCHAR(255),
		gender VARCHAR(10),
		date_of_birth TIMESTAMP,
		email VARCHAR(255) UNIQUE,
		password VARCHAR(255),
		role VARCHAR(20)
	);`
	_, err := uc.db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

func (uc *UserController) CreateUser(u *models.User) uuid.UUID {
	err := uc.db.QueryRow(
		`INSERT INTO users (id, first_name, last_name, gender, date_of_birth, email, password, role)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`, u.ID, u.FirstName, u.LastName, u.Gender, u.DateOfBirth, u.Email, u.Password, u.Role).Scan(&u)
	if err != nil {
		log.Fatal(err)
	}
	return u.ID
}

func (uc *UserController) UpdateUser(u *models.User) {
	_, err := uc.db.Exec(
		`UPDATE users
		SET first_name = $2, last_name = $3, gender = $4, date_of_birth = $5, email = $6, password = $7, role = $8
		WHERE id = $1`, u.ID, u.FirstName, u.LastName, u.Gender, u.DateOfBirth, u.Email, u.Password, u.Role)
	if err != nil {
		log.Fatal(err)
	}
}

func (uc *UserController) GetUserByEmail(email string) models.User {
	var user models.User
	err := uc.db.QueryRow(
		`SELECT id, first_name, last_name, gender, date_of_birth, email, password, role
		FROM users WHERE email = $1`, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Gender, &user.DateOfBirth, &user.Email, &user.Password, &user.Role)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func (uc *UserController) GetUsersByRole(role string) []models.User {
	var users []models.User

	rows, err := uc.db.Query("SELECT id, first_name, last_name, gender, date_of_birth, email, password, role FROM users WHERE role = $1", role)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Gender, &user.DateOfBirth, &user.Email, &user.Password, &user.Role)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users
}
