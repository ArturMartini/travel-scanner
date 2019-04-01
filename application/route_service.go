package application

import "domain"

type RouteService struct {
	RouteRepo domain.RouteRepository
}

func (s *RouteService) All() ([]domain.Route, error) {
	return s.RouteRepo.All()
}

func (s *RouteService) SaveRoute(route *domain.Route) error {
	route := domain.Route {

	}

	if err := s.RouteRepo.Save(&route); err != nil {
		return err
	}
}

func (s *RouteService) FindBestRoute() ([]domain.Route, error) {
	
}
