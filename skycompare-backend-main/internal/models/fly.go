package models

type Fly struct {
	Airway  Route   `json:"route"`
	Corpor  Company `json:"company"`
	HourDep string  `json:"timeDep"`
	HourArr string  `json:"timeArr"`
	Price   float32 `json:"price"`
	Sales   int32   `json:"sales"`
}
