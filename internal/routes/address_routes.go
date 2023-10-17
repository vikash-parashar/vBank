package routes

import (
	"vbank/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterAddressRoutes(r *mux.Router) {
	r.HandleFunc("/addresses", handlers.ListAddresses).Methods("GET")
	r.HandleFunc("/addresses/{id}", handlers.GetAddress).Methods("GET")
	r.HandleFunc("/addresses", handlers.CreateAddress).Methods("POST")
	r.HandleFunc("/addresses/{id}", handlers.UpdateAddress).Methods("PUT")
	r.HandleFunc("/addresses/{id}", handlers.DeleteAddress).Methods("DELETE")
}
