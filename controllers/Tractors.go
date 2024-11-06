package controllers

import (
	"NLPF-Assets/database"
	"NLPF-Assets/models"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// {
// 	"client_id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
// 	"tractor_name": "string",
// 	"volume": 0,
// 	"type": "bulk",
// 	"min_price": 100,
// 	"start_checkpoint_id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
// 	"end_checkpoint_id": "3fa85f64-5717-4562-b3fc-2c963f66afa6"
//   }
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