package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Modèle pour un post
type Post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// Connexion à MongoDB
func connectDB() {
	uri := "mongodb://ligne8:ligne8password@localhost:27017" // ou via variable d'environnement MONGO_URI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	// Vérifier la connexion
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

// Gestionnaire pour la route POST
func createPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var post Post

	// Décoder le corps de la requête JSON
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Connexion à la collection "posts"
	collection := client.Database("nlpf_assets").Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Insérer le post dans MongoDB
	result, err := collection.InsertOne(ctx, bson.M{
		"title": post.Title,
		"body":  post.Body,
	})
	if err != nil {
		http.Error(w, "Error inserting document", http.StatusInternalServerError)
		return
	}

	// Répondre avec l'ID du document inséré
	fmt.Fprintf(w, "Inserted document with ID: %v", result.InsertedID)
}

func main() {
	connectDB()

	http.HandleFunc("/posts", createPost)

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
