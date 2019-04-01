package application

import (
	"fmt"
	"github.com/hdiomede/travel-scanner/domain"
)

type RouteService struct {
	RouteRepo domain.RouteRepository
}

func (s *RouteService) All() ([]domain.Route, error) {
	return s.RouteRepo.All()
}

func (s *RouteService) SaveRoute(route *domain.Route) error {
	if s.RouteRepo.Exists(route) {
		fmt.Println("Ja existe")
	}

	s.RouteRepo.Save(route)

	return nil
}

func (s *RouteService) FindBestRoute() ([]domain.Route, error) {
	return s.RouteRepo.All()
}
