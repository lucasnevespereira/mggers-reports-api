package models

import "time"

type Report struct {
	Id         string    `json:"id"`
	Position   Position  `json:"position"`
	ReportedAt time.Time `json:"reportedAt"`
}

type Position struct {
	Latitude  int `json:"latitude"`
	Longitude int `json:"longitude"`
}

type CreateReport struct {
	Position   Position  `json:"position"`
}