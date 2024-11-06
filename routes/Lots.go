package routes

import (
	"NLPF-Assets/controllers"
	"net/http"
)

func LotRoutes(){
	http.HandleFunc("/lots", controllers.CreateLot)
}