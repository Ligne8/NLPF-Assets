package main

import (
	"NLPF-Assets/database"
	"NLPF-Assets/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.ConnectDB()
	r := mux.NewRouter()
    http.Handle("/", r)
	routes.LotRoutes(r)
	routes.TractorRoutes(r)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
