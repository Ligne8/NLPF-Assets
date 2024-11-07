package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    databaseURL := os.Getenv("DATABASE_URL")
    if databaseURL == "" {
        log.Fatal("DATABASE_URL is not set in the environment")
    }
    clientOptions := options.Client().ApplyURI(databaseURL)
    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    Client = client
	fmt.Println("Connected to MongoDB!")
}