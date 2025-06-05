package models

type User struct {
	ID         int    `json:"id"`
	Nickname   string `json:"nickname"`
	FullName   string `json:"full_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	FavAirport string `json:"fav_airport"`
}
