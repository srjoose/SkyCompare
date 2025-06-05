package service

import (
	"errors"
	"math/rand"
	"skycompare-backend-main/internal/models"
	"testing"
)

type MockRouteRepo struct {
	GetRouteFunc     func(dep, arr string) (*models.Route, error)
	GetCompaniesFunc func() ([]models.Company, error)
}

func (m *MockRouteRepo) GetRoute(dep, arr string) (*models.Route, error) {
	return m.GetRouteFunc(dep, arr)
}

func (m *MockRouteRepo) GetCompanies() ([]models.Company, error) {
	return m.GetCompaniesFunc()
}

func TestGetRoute_Success(t *testing.T) {
	mockRepo := &MockRouteRepo{
		GetRouteFunc: func(dep, arr string) (*models.Route, error) {
			return &models.Route{
				ID:          1,
				Departure:   dep,
				Destination: arr,
				Duration:    120,
				AvgPrice:    100.0,
			}, nil
		},
		GetCompaniesFunc: func() ([]models.Company, error) {
			return []models.Company{
				{ID: 1, Name: "Company A", Multiply: 1.0},
				{ID: 2, Name: "Company B", Multiply: 1.1},
			}, nil
		},
	}

	// For reproducibility
	rand.Seed(42)

	service := RouteService{Repo: mockRepo}
	flies, err := service.GetRoute("AAA", "BBB")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(flies) == 0 {
		t.Errorf("Expected at least 1 fly, got 0")
	}
}

func TestGetRoute_Fail(t *testing.T) {
	mockRepo := &MockRouteRepo{
		GetRouteFunc: func(dep, arr string) (*models.Route, error) {
			return nil, errors.New("Route not found")
		},
		GetCompaniesFunc: func() ([]models.Company, error) {
			return nil, nil
		},
	}

	service := RouteService{Repo: mockRepo}
	_, err := service.GetRoute("AAA", "BBB")

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
