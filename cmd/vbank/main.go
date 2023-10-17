package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"vbank/internal/db"
	"vbank/internal/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables from the .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Read environment variables and set them in the config struct
	db.AppConfig.DBHost = os.Getenv("DB_HOST")
	db.AppConfig.DBPort = os.Getenv("DB_PORT")
	db.AppConfig.DBName = os.Getenv("DB_NAME")
	db.AppConfig.DBUser = os.Getenv("DB_USER")
	db.AppConfig.DBPassword = os.Getenv("DB_PASSWORD")
	db.AppConfig.ServerPort = os.Getenv("PORT")
	db.AppConfig.JWTSecret = os.Getenv("JWT_SECRET")
	db.AppConfig.JWTExpiration = os.Getenv("JWT_EXPIRATION")
	db.AppConfig.DebugMode = os.Getenv("DEBUG_MODE")
	db.AppConfig.NodeEnv = os.Getenv("NODE_ENV")
}

func main() {

	// Initialize the database connection
	db.InitDB()
	defer db.GetDB().Close() // Ensure the database connection is closed when the application exits
	// Create a new Gorilla Mux router
	r := mux.NewRouter()

	// Register routes for different components
	routes.RegisterUserRoutes(r)
	routes.RegisterTransactionRoutes(r)
	routes.RegisterNomineeRoutes(r)
	routes.RegisterAddressRoutes(r)
	routes.RegisterAccountRoutes(r)

	// Start your server
	serverAddr := fmt.Sprintf(":%s", db.AppConfig.ServerPort)
	log.Printf("Server is running on port %s", db.AppConfig.ServerPort)
	if err := http.ListenAndServe(serverAddr, r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
