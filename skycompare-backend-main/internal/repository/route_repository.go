package repository

import (
	"database/sql"
	"skycompare-backend-main/internal/models"
)


type RouteRepositoryInterface interface {
	GetRoute(dep, arr string) (*models.Route, error)
	GetCompanies() ([]models.Company, error)
}

type RouteRepository struct {
	DB *sql.DB
}

func (r *RouteRepository) GetRoute(dep, arr string) (*models.Route, error) {
	query := "SELECT id, departure, destination, duration, avg_price FROM routes WHERE departure=? AND destination=?"
	row := r.DB.QueryRow(query, dep, arr)

	var route models.Route
	err := row.Scan(&route.ID, &route.Departure, &route.Destination, &route.Duration, &route.AvgPrice)
	if err != nil {
		return nil, err
	}
	return &route, nil
}

func (r *RouteRepository) GetCompanies() ([]models.Company, error) {
	query := "SELECT id, name, multiply FROM company"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []models.Company
	for rows.Next() {
		var c models.Company
		err := rows.Scan(&c.ID, &c.Name, &c.Multiply)
		if err != nil {
			return nil, err
		}
		companies = append(companies, c)
	}
	return companies, nil
}
