package repository

import (
	"database/sql"
	"skycompare-backend-main/internal/models"
)


type AirportRepositoryInterface interface {
	GetAllAirports() ([]models.Airport, error)
}

type AirportRepository struct {
	DB *sql.DB
}

func (r *AirportRepository) GetAllAirports() ([]models.Airport, error) {
	query := "SELECT id, name, location, IATA FROM airports"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var airports []models.Airport
	for rows.Next() {
		var a models.Airport
		err := rows.Scan(&a.ID, &a.Name, &a.Location, &a.IATA)
		if err != nil {
			return nil, err
		}
		airports = append(airports, a)
	}
	return airports, nil
}
