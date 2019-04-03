package application

import (
	"github.com/hdiomede/travel-scanner/domain"
	"github.com/hdiomede/travel-scanner/errors"
)

type FlightService struct {
	FlightRepo domain.FlightRepository
	bookingService BookingService
	flights  domain.Flights
}


func NewFlightService(flightRepository domain.FlightRepository) *FlightService {
	var flights = domain.Flights{make(map[string]map[string]int)}
	flightService := FlightService{FlightRepo: flightRepository, flights: flights, bookingService: BookingService{&flights}}
	flightService.LoadFlights()

	return &flightService
}

func (flightService *FlightService) LoadFlights() {
	flightsList, _ := flightService.FlightRepo.All()

	for _, flight := range flightsList {
		flightService.flights.AddFlight(&flight)
	}
}

func (flightService *FlightService) All() ([]domain.Flight, error) {
	return flightService.FlightRepo.All()
}

func (flightService *FlightService) SaveFlight(flight *domain.Flight) error {
	if flightService.FlightRepo.Exists(flight) {
		return errors.FlightAlreadyExists()
	}

	flightService.FlightRepo.Save(flight)
	flightService.flights.AddFlight(flight)

	return nil
}

func (flightService *FlightService) FindBestFlight(origin string, dest string) error {
	return flightService.bookingService.FindBestFlight(origin, dest)
}
