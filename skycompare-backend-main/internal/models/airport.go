package models

type Airport struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	IATA     string `json:"IATA"`
}
