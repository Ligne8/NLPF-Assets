package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tractor struct {
	Id          			primitive.ObjectID  `bson:"_id,omitempty" json:"_id" `
	ClientId 				string              `bson:"client_id,omitempty" json:"client_id"`	
	TractorName 			string			  	`bson:"name,omitempty" json:"name"`
	Status 					Status              `bson:"status,omitempty" json:"status"`
	Volume 					float64             `bson:"volume,omitempty" json:"volume"`
	OccupiedVolume 			float64				`bson:"occupied_volume,omitempty" json:"occupied_volume"`
	RouteId 				string              `bson:"route_id,omitempty" json:"route_id"`
	Type 					Type                `bson:"type,omitempty" json:"type"`
	MinPrice 				float64             `bson:"min_price,omitempty" json:"min_price"`
    CurrentCheckpointId     string              `bson:"current_checkpoint,omitempty" json:"current_checkpoint"`
    StartCheckpointId       string              `bson:"start_checkpoint,omitempty" json:"start_checkpoint"`
    EndCheckpointId         string              `bson:"end_checkpoint,omitempty" json:"end_checkpoint"`
}