package main

import (
	"NLPF-Assets/database"
	"NLPF-Assets/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	database.ConnectDB()
	routes.LotRoutes()

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
