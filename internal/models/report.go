package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Report struct {
	ID          primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty`
	Description string             `json:"description, omitempty" bson:"description, omitempty`
	Position    Position           `json:"position, omitempty" bson:"position, omitempty`
	ReportedAt  string             `json:"reportedAt, omitempty" bson:"reported_at, omitempty`
}

type ReportRequest struct {
	Description string   `json:"description"`
	Position    Position `json:"position"`
	ReportedAt  string   `json:"reportedAt"`
}

type Position struct {
	Latitude  float64 `json:"latitude, omitempty" bson:"latitude, omitempty`
	Longitude float64 `json:"longitude, omitempty" bson:"longitude, omitempty`
}
