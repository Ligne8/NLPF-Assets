package routes

import (
	"NLPF-Assets/controllers"

	"github.com/gorilla/mux"
)


func TractorRoutes(r *mux.Router) {
	r.HandleFunc("/tractors", controllers.CreateTractor).Methods("POST")

}