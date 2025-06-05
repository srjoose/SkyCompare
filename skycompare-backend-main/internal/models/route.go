package models

type Route struct {
	ID          int     `json:"id"`
	Departure   string  `json:"departure"`
	Destination string  `json:"arrival"`
	Duration    int     `json:"duration"`
	AvgPrice    float32 `json:"price"`
}
