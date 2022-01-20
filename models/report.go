package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Report struct {
	ID          primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty`
	Description string             `json:"description, omitempty" bson:"description, omitempty`
	Latitude    int                `json:"latitude, omitempty" bson:"latitude, omitempty`
	Longitude   int                `json:"longitude, omitempty" bson:"longitude, omitempty`
	ReportedAt  time.Time          `json:"reported_at, omitempty" bson:"reported_at, omitempty`
}

type ReportRequest struct {
	Description string    `json:"description"`
	Latitude    int       `json:"latitude"`
	Longitude   int       `json:"longitude"`
	ReportedAt  time.Time `json:"reported_at"`
}
