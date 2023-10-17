package routes

import (
	"vbank/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterNomineeRoutes(r *mux.Router) {
	r.HandleFunc("/nominees", handlers.ListNominee).Methods("GET")
	r.HandleFunc("/nominees/{id}", handlers.GetNominee).Methods("GET")
	r.HandleFunc("/nominees", handlers.CreateNominee).Methods("POST")
	r.HandleFunc("/nominees/{id}", handlers.UpdateNominee).Methods("PUT")
	r.HandleFunc("/nominees/{id}", handlers.DeleteNominee).Methods("DELETE")
}
