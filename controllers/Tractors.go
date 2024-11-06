package controllers

import (
	"NLPF-Assets/database"
	"NLPF-Assets/models"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTractorRequest struct {
	ClientId          string  		`json:"client_id"`
	TractorName       string  		`json:"tractor_name"`
	Volume            float64 		`json:"volume"`
	Type              models.Type  	`json:"type"`
	MinPrice          float64 		`json:"min_price"`
	StartCheckpointId string  		`json:"start_checkpoint_id"`
	EndCheckpointId   string  		`json:"end_checkpoint_id"`
}

func CreateTractor(w http.ResponseWriter, r *http.Request) {
    var req CreateTractorRequest
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
    if req.MinPrice <= 0 {
        http.Error(w, "min_price must be greater than 0", http.StatusBadRequest)
        return
    }
    if req.StartCheckpointId == "" {
        http.Error(w, "start_checkpoint_id is required", http.StatusBadRequest)
        return
    }
    if req.EndCheckpointId == "" {
        http.Error(w, "end_checkpoint_id is required", http.StatusBadRequest)
        return
    }

    tractor := models.Tractor{
        Id:                primitive.NewObjectID(),
        ClientId:          req.ClientId,
        TractorName:       req.TractorName,
        Volume:            req.Volume,
        Type:              req.Type,
        MinPrice:          req.MinPrice,
        StartCheckpointId: req.StartCheckpointId,
        EndCheckpointId:   req.EndCheckpointId,
    }

    collection := database.Client.Database("assets").Collection("tractors")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := collection.InsertOne(ctx, tractor)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tractor)
}

func GetAllTractorsByClient(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    clientId := vars["client_id"]

    collection := database.Client.Database("assets").Collection("tractors")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, bson.M{"client_id": clientId})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var tractors []models.Tractor
    if err = cursor.All(ctx, &tractors); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tractors)
}

func GetTractorById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    tractorId := vars["tractor_id"]

    objectId, err := primitive.ObjectIDFromHex(tractorId)
    if err != nil {
        http.Error(w, "Invalid tractor ID", http.StatusBadRequest)
        return
    }

    collection := database.Client.Database("assets").Collection("tractors")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var tractor models.Tractor
    err = collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&tractor)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tractor)
}

func UpdateTractorStatus(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    tractorId := vars["tractor_id"]

    objectId, err := primitive.ObjectIDFromHex(tractorId)
    if err != nil {
        http.Error(w, "Invalid tractor ID", http.StatusBadRequest)
        return
    }

    var req struct {
        Status models.Status `json:"status"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    collection := database.Client.Database("assets").Collection("tractors")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err = collection.UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": bson.M{"status": req.Status}})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func GetAllTractorsInTransit(w http.ResponseWriter, r *http.Request) {
    collection := database.Client.Database("assets").Collection("tractors")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, bson.M{"status": string(models.StatusInTransit)})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var tractors []models.Tractor
    if err = cursor.All(ctx, &tractors); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tractors)
}

func UpdateTractorCurrentCheckpoint(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    tractorId := vars["tractor_id"]

    objectId, err := primitive.ObjectIDFromHex(tractorId)
    if err != nil {
        http.Error(w, "Invalid tractor ID", http.StatusBadRequest)
        return
    }

    var req struct {
        CheckpointId string `json:"checkpoint_id"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    collection := database.Client.Database("assets").Collection("tractors")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err = collection.UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": bson.M{"current_checkpoint_id": req.CheckpointId}})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}