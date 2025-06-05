package service

import (
	"skycompare-backend-main/internal/models"
	"skycompare-backend-main/internal/repository"
)

type AirportService struct {
	Repo repository.AirportRepositoryInterface
}


func (s *AirportService) GetAllAirports() ([]models.Airport, error) {
	return s.Repo.GetAllAirports()
}

func (s *AirportService) GetAirportsWithoutOne(exclude string) ([]models.Airport, error) {
	all, err := s.Repo.GetAllAirports()
	if err != nil {
		return nil, err
	}

	filtered := []models.Airport{}
	for _, airport := range all {
		if airport.IATA != exclude {
			filtered = append(filtered, airport)
		}
	}
	return filtered, nil
}
