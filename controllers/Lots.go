package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"NLPF-Assets/database"
	"NLPF-Assets/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateLotRequest struct {
    ClientId          string  `json:"client_id"`
    Volume            float64 `json:"volume"`
    Type              string  `json:"type"`
    MaxPrice          float64 `json:"max_price"`
    StartCheckpointId string  `json:"start_checkpoint_id"`
    EndCheckpointId   string  `json:"end_checkpoint_id"`
}

func GetLotsByClient(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    clientId := vars["client_id"]

    collection := database.Client.Database("assets").Collection("lots")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    filter := bson.M{"client_id": clientId}
    cursor, err := collection.Find(ctx, filter)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx)

    var lots []models.Lot
    if err = cursor.All(ctx, &lots); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(lots)
}

func CreateLot(w http.ResponseWriter, r *http.Request) {
    var req CreateLotRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	if req.ClientId == "" {
		http.Error(w, "client_id is required", http.StatusBadRequest)
		return
	}
	if req.Volume <= 0 {
		http.Error(w, "volume must be greater than 0", http.StatusBadRequest)
		return
	}
	if req.Type == "" {
		http.Error(w, "type is required", http.StatusBadRequest)
		return
	}
	if req.MaxPrice <= 0 {
		http.Error(w, "max_price must be greater than 0", http.StatusBadRequest)
		return
	}
	//TODO: Check if start_checkpoint_id exists
	if req.StartCheckpointId == "" {
		http.Error(w, "start_checkpoint_id is required", http.StatusBadRequest)
		return
	}
	//TODO: Check if end_checkpoint_id exists
	if req.EndCheckpointId == "" {
		http.Error(w, "end_checkpoint_id is required", http.StatusBadRequest)
		return
	}


    lot := models.Lot{
        Id:                primitive.NewObjectID(),
        ClientId:          req.ClientId,
        Volume:            req.Volume,
        Type:              req.Type,
        MaxPrice:          req.MaxPrice,
        StartCheckpointId: req.StartCheckpointId,
        EndCheckpointId:   req.EndCheckpointId,
    }

    collection := database.Client.Database("assets").Collection("lots")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := collection.InsertOne(ctx, lot)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(lot)
}