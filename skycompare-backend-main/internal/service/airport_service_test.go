package service

import (
	"errors"
	"skycompare-backend-main/internal/models"
	"testing"
)

type MockAirportRepo struct {
	GetAllAirportsFunc func() ([]models.Airport, error)
}

func (m *MockAirportRepo) GetAllAirports() ([]models.Airport, error) {
	return m.GetAllAirportsFunc()
}

func TestGetAllAirports_Success(t *testing.T) {
	mockRepo := &MockAirportRepo{
		GetAllAirportsFunc: func() ([]models.Airport, error) {
			return []models.Airport{
				{ID: 1, Name: "Airport A", Location: "Location A", IATA: "AAA"},
				{ID: 2, Name: "Airport B", Location: "Location B", IATA: "BBB"},
			}, nil
		},
	}

	service := AirportService{Repo: mockRepo}
	airports, err := service.GetAllAirports()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(airports) != 2 {
		t.Errorf("Expected 2 airports, got %d", len(airports))
	}
}

func TestGetAllAirports_Error(t *testing.T) {
	mockRepo := &MockAirportRepo{
		GetAllAirportsFunc: func() ([]models.Airport, error) {
			return nil, errors.New("DB error")
		},
	}

	service := AirportService{Repo: mockRepo}
	_, err := service.GetAllAirports()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
