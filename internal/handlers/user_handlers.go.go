// user/handlers.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"vbank/internal/auth"
	"vbank/internal/controllers"
	"vbank/internal/models"

	"github.com/google/uuid"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Parse user registration data from the request body
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// Check if a user with the provided email already exists
	userController := controllers.NewUserController()
	_, err := userController.GetUserByEmail(user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			// User with the provided email doesn't exist, proceed with registration
			newUser := models.User{
				ID:          uuid.New(),
				FirstName:   user.FirstName,
				LastName:    user.LastName,
				Gender:      user.Gender,
				DateOfBirth: user.DateOfBirth,
				Email:       user.Email,
				Password:    user.Password, // Remember to hash the password
				Role:        user.Role,
			}

			_, registrationErr := userController.CreateUser(&newUser)
			if registrationErr != nil {
				log.Println(registrationErr)
				http.Error(w, "Failed to register user", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
		} else {
			log.Println(err)
			http.Error(w, "Failed to check for existing user", http.StatusInternalServerError)
		}
	} else {
		// User with the provided email already exists
		http.Error(w, "User with this email already exists", http.StatusConflict)
	}
}

// LoginUser handles user login and generates a JWT token.
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// Retrieve the user from the database by email
	userController := controllers.NewUserController()
	user, err := userController.GetUserByEmail(loginData.Email)
	if err != nil {
		log.Println(err)
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Verify the hashed password
	if !auth.VerifyPassword(loginData.Password, user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate a JWT token
	token, err := auth.GenerateToken(&user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Return the token in the response
	response := map[string]string{"token": token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {

}

func ListUsers(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
