package routes

import (
	"go-api/controllers/postcontroller"

	"github.com/gorilla/mux"
)

func PostRoutes(r *mux.Router) {
	router := r.PathPrefix("/posts").Subrouter()

	router.HandleFunc("", postcontroller.Index).Methods("GET")
	router.HandleFunc("/{slug}", postcontroller.Show).Methods("GET")
}
