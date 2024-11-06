package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status string

const (
    StatusPending        	Status = "pending"
	StatusAtTrader	   		Status = "at_trader"
	StatusOnMarket	   		Status = "on_market"
	StatusArchive	   		Status = "archive"
	StatusAvailable	   		Status = "available"
	StatusInTransit 		Status = "in_transit"
	StatusReturnFromMarket 	Status = "return_from_market"
)

type Type string

const (
    TypeBulk        	Type = "bulk"
    TypeContainer   	Type = "container"
    TypeSolid      	    Type = "solid"
)

type Lot struct {
    Id                      primitive.ObjectID 	`bson:"id,omitempty" json:"id"`
    ClientId                string              `bson:"client_id,omitempty" json:"client_id"`
    Status                  Status        	    `bson:"status,omitempty" json:"status"`
    Volume                  float64            	`bson:"volume,omitempty" json:"volume"`
    CreatedAt               string              `bson:"created_at,omitempty" json:"created_at"`
    Type                    string              `bson:"type,omitempty" json:"type"`
    MaxPrice                float64             `bson:"max_price,omitempty" json:"max_price"`
    CurrentCheckpointId     string              `bson:"current_checkpoint,omitempty" json:"current_checkpoint"`
    StartCheckpointId       string              `bson:"start_checkpoint,omitempty" json:"start_checkpoint"`
    EndCheckpointId         string              `bson:"end_checkpoint,omitempty" json:"end_checkpoint"`
}