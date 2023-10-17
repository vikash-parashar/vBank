package routes

import (
	"vbank/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterAccountRoutes(r *mux.Router) {
	r.HandleFunc("/accounts", handlers.ListAccounts).Methods("GET")
	r.HandleFunc("/accounts/{id}", handlers.GetAccount).Methods("GET")
	r.HandleFunc("/accounts", handlers.CreateAccount).Methods("POST")
	r.HandleFunc("/accounts/{id}", handlers.UpdateAccount).Methods("PUT")
	r.HandleFunc("/accounts/{id}", handlers.DeleteAccount).Methods("DELETE")
}
