package application

import (
	"fmt"
	"errors"
	"github.com/hdiomede/travel-scanner/domain"
)

type RouteService struct {
	RouteRepo domain.RouteRepository
	routeMap  map[string]map[string]int
}

func NewRouteService(routeRepository domain.RouteRepository) *RouteService {
	routeService := RouteService{RouteRepo: routeRepository, routeMap: make(map[string]map[string]int)}
	routeService.LoadRoutes()

	return &routeService
}

func (s *RouteService) LoadRoutes() {
	routesList, _ := s.RouteRepo.All()
	s.buildMatrix(routesList)
}

func (s *RouteService) All() ([]domain.Route, error) {
	return s.RouteRepo.All()
}

func (s *RouteService) SaveRoute(route *domain.Route) error {
	if s.RouteRepo.Exists(route) {
		return errors.New("Route already exists")
	}

	s.RouteRepo.Save(route)
	s.addRouteToMatrix(route)

	return nil
}

func (s *RouteService) FindBestRoute() ([]domain.Route, error) {
	return s.RouteRepo.All()
}

func (s *RouteService) buildMatrix(routes []domain.Route) {
	for _, r := range routes {
		s.addRouteToMatrix(&r)
	}
}

func (s *RouteService) addRouteToMatrix(route *domain.Route) {
	child, ok := s.routeMap[route.From]
	if !ok {
		child = map[string]int{}
		s.routeMap[route.From] = child
	}

	child[route.To] = route.Price

	fmt.Println(route.To)
}

func (s *RouteService) PrintMatrixElement(origin string, target string) {
	fmt.Println(s.routeMap[origin][target])
}
