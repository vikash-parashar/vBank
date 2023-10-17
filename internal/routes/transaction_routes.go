package routes

import (
	"vbank/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterTransactionRoutes(r *mux.Router) {
	r.HandleFunc("/transactions", handlers.ListTransactions).Methods("GET")
	r.HandleFunc("/transactions/{id}", handlers.GetTransaction).Methods("GET")
	r.HandleFunc("/transactions", handlers.CreateTransaction).Methods("POST")
	r.HandleFunc("/transactions/{id}", handlers.UpdateTransaction).Methods("PUT")
	r.HandleFunc("/transactions/{id}", handlers.DeleteTransaction).Methods("DELETE")
}
