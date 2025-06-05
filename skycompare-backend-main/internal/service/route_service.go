package service

import (
	"math/rand"
	"strconv"
	"skycompare-backend-main/internal/models"
	"skycompare-backend-main/internal/repository"
)

type RouteService struct {
	Repo repository.RouteRepositoryInterface
}

func (s *RouteService) GetRoute(dep, arr string) ([]models.Fly, error) {
	route, err := s.Repo.GetRoute(dep, arr)
	if err != nil {
		return nil, err
	}

	companies, err := s.Repo.GetCompanies()
	if err != nil {
		return nil, err
	}

	var airways []models.Fly
	rnd := rand.Intn(4)
	comps := map[int]bool{}

	for count := 0; count <= rnd; count++ {
		num := rand.Intn(len(companies))
		if comps[num] {
			count--
			continue
		}
		comps[num] = true

		ranComponent := rand.Float32()*(1.35-0.65) + 0.65
		price := route.AvgPrice * companies[num].Multiply * ranComponent
		var sales int32
		if ranComponent <= 0.8 {
			sales = 1
		}

		timeDep := rand.Float32() * 23.99
		minDep := timeDep * 60
		minArr := minDep + float32(route.Duration)
		if minArr > 1440 {
			minArr -= 1440
		}

		hourDep := formatTime(minDep)
		hourArr := formatTime(minArr)

		f := models.Fly{*route, companies[num], hourDep, hourArr, price, sales}
		airways = append(airways, f)
	}

	return airways, nil
}

func formatTime(minutes float32) string {
	hour := int(minutes) / 60
	min := int(minutes) % 60
	return strconv.Itoa(hour) + ":" + twoDigits(min)
}

func twoDigits(val int) string {
	if val < 10 {
		return "0" + strconv.Itoa(val)
	}
	return strconv.Itoa(val)
}
