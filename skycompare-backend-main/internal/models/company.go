package models

type Company struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Multiply float32 `json:"multiply"`
}
