package routes

import (
	"NLPF-Assets/controllers"

	"github.com/gorilla/mux"
)


func TractorRoutes(r *mux.Router) {
	r.HandleFunc("/tractors", controllers.CreateTractor).Methods("POST")
	r.HandleFunc("/tractors/clients/{client_id}", controllers.GetAllTractorsByClient).Methods("GET")
	r.HandleFunc("/tractors/{tractor_id}", controllers.GetTractorById).Methods("GET")
	r.HandleFunc("/tractors/{tractor_id}/status", controllers.UpdateTractorStatus).Methods("PUT")
	r.HandleFunc("/tractors/status/in_transit", controllers.GetAllTractorsInTransit).Methods("GET")
	r.HandleFunc("/tractors/{tractor_id}/current_checkpoint", controllers.UpdateTractorCurrentCheckpoint).Methods("PUT")
}