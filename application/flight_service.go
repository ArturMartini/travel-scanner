package application

import (
	"errors"
	"github.com/hdiomede/travel-scanner/domain"
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
		return errors.New("Flight already exists")
	}

	flightService.FlightRepo.Save(flight)
	flightService.flights.AddFlight(flight)

	return nil
}

func (flightService *FlightService) FindBestFlight(origin string, dest string) {
	flightService.bookingService.FindBestFlight(origin, dest)
}
