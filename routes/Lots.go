package routes

import (
    "NLPF-Assets/controllers"
    "github.com/gorilla/mux"
)

func LotRoutes(r *mux.Router) {
    r.HandleFunc("/lots", controllers.CreateLot).Methods("POST")
    r.HandleFunc("/lots/clients/{client_id}", controllers.GetLotsByClient).Methods("GET")
	r.HandleFunc("/lots/{lot_id}", controllers.GetLotById).Methods("GET")
	r.HandleFunc("/lots/{lot_id}/status", controllers.UpdateLotStatus).Methods("PUT")
	r.HandleFunc("/lots/{lot_id}/current_checkpoint", controllers.UpdateLotCurrentCheckpoint).Methods("PUT")

}