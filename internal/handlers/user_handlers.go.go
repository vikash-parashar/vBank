// user/handlers.go
package handlers

import (
	"net/http"
	"vbank/internal/auth"
	"vbank/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

}

// LoginUser handles user login and generates a JWT token.
func LoginUser(c *gin.Context) {
	// Parse user login credentials from the request body
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Retrieve the user from the database by email
	userController := controllers.NewUserController()
	user := userController.GetUserByEmail(loginData.Email)

	// Verify the hashed password
	if !auth.VerifyPassword(loginData.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate a JWT token
	token, err := auth.GenerateToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return the token in the response
	c.JSON(http.StatusOK, gin.H{"token": token})
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
