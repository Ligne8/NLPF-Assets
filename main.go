package main

import (
	"NLPF-Assets/database"
	"NLPF-Assets/routes"
	"log"
	"net/http"
)

func main() {
	database.ConnectDB()
	routes.LotRoutes()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
