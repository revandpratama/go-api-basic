package routes

import (
	"go-api/controllers/usercontroller"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	router := r.PathPrefix("/users").Subrouter()

	router.HandleFunc("", usercontroller.Index).Methods("GET")
	router.HandleFunc("", usercontroller.Create).Methods("POST")
	router.HandleFunc("/{username}", usercontroller.Show).Methods("GET")

}
