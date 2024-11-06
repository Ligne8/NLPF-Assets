package routes

import (
    "NLPF-Assets/controllers"
    "github.com/gorilla/mux"
    "net/http"
)

func LotRoutes() {
    r := mux.NewRouter()
    r.HandleFunc("/lots", controllers.CreateLot).Methods("POST")
    r.HandleFunc("/lots/clients/{client_id}", controllers.GetLotsByClient).Methods("GET")

    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}