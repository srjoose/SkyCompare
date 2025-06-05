package routes

import (
	"database/sql"
	"net/http"
	"skycompare-backend-main/internal/handler"
	"skycompare-backend-main/internal/repository"
	"skycompare-backend-main/internal/service"
)

func SetupRoutes(db *sql.DB) http.Handler {
	userRepo := &repository.UserRepository{DB: db}
	userService := &service.UserService{Repo: userRepo}
	userHandler := &handler.UserHandler{Service: userService}

	airportRepo := &repository.AirportRepository{DB: db}
	airportService := &service.AirportService{Repo: airportRepo}
	airportHandler := &handler.AirportHandler{Service: airportService}

	routeRepo := &repository.RouteRepository{DB: db}
	routeService := &service.RouteService{Repo: routeRepo}
	routeHandler := &handler.RouteHandler{Service: routeService}

	mux := http.NewServeMux()

	mux.HandleFunc("/register", userHandler.Register)
	mux.HandleFunc("/login", userHandler.Login)
	mux.HandleFunc("/airports", airportHandler.GetAll)
	mux.HandleFunc("/airports/selected", airportHandler.GetWithoutOne)
	mux.HandleFunc("/routes", routeHandler.GetRoutes)
	mux.HandleFunc("/favAirport", userHandler.UpdateFavourite)
	mux.HandleFunc("/favAirportGet", userHandler.GetFavourite)

	return mux
}
