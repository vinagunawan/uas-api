package router

import (
	"backend/service/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/water", handlers.RecordWaterIntake).Methods("POST")
	r.HandleFunc("/water/{userId:[0-9]+}", handlers.GetWaterIntake).Methods("GET")
}
