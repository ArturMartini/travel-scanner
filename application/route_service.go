package application

import (
	"errors"
	"github.com/hdiomede/travel-scanner/domain"
)

type RouteService struct {
	RouteRepo domain.RouteRepository
	routeMap  domain.RouteMap
}


func NewRouteService(routeRepository domain.RouteRepository) *RouteService {
	routeService := RouteService{RouteRepo: routeRepository, routeMap: domain.RouteMap{make(map[string]map[string]int)}}
	routeService.LoadRoutes()

	return &routeService
}

func (s *RouteService) LoadRoutes() {
	routesList, _ := s.RouteRepo.All()
	s.routeMap.BuildMatrix(routesList)
}

func (s *RouteService) All() ([]domain.Route, error) {
	return s.RouteRepo.All()
}

func (s *RouteService) SaveRoute(route *domain.Route) error {
	if s.RouteRepo.Exists(route) {
		return errors.New("Route already exists")
	}

	s.RouteRepo.Save(route)
	s.routeMap.AddRoute(route)

	return nil
}

func (s *RouteService) FindBestRoute() {
	s.routeMap.FindBestRoute("GRU", "")
}
