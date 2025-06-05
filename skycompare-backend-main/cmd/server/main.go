package main

import (
	"log"
	"net/http"
	"skycompare-backend-main/internal/config"
	"skycompare-backend-main/internal/database"
	"skycompare-backend-main/internal/middleware"
	"skycompare-backend-main/internal/routes"
)

func main() {
	config.LoadEnv()

	db := database.Connect()
	defer db.Close()

	router := routes.SetupRoutes(db)
	log.Println("Servidor escuchando en :5152")
	http.ListenAndServe(":5152", middleware.Cors(router))
}
