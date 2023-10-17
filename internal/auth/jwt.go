package auth

import (
	"time"
	"vbank/internal/models"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("your-secret-key")

// GenerateToken generates a JWT token for a user with specified claims.
func GenerateToken(u *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = u.ID // Convert UUID to string
	claims["email"] = u.Email
	claims["role"] = u.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time (24 hours)
	return token.SignedString(jwtSecret)
}

// ParseToken parses and validates a JWT token and returns the claims.
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
