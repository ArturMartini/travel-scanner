package application

import (
	"github.com/hdiomede/travel-scanner/domain"
	"github.com/hdiomede/travel-scanner/errors"
)

type FlightService interface {
	All() ([]domain.Flight, error)
	FindBestFlight(flight domain.Flight) error
	SaveFlight(flight *domain.Flight) error
}

type flightService struct {
	FlightRepo domain.FlightRepository
	bookingService BookingService
	flights  domain.Flights
}

func NewFlightService(flightRepository domain.FlightRepository) *flightService {
	var flights = domain.Flights{make(map[string]map[string]int)}
	fs := flightService{FlightRepo: flightRepository, flights: flights, bookingService: BookingService{&flights}}
	fs.loadFlights()

	return &fs
}

func (fs *flightService) loadFlights() {
	flightsList, _ := fs.FlightRepo.All()

	for _, flight := range flightsList {
		fs.flights.AddFlight(&flight)
	}
}

func (fs *flightService) All() ([]domain.Flight, error) {
	return fs.FlightRepo.All()
}

func (fs *flightService) SaveFlight(flight *domain.Flight) error {
	if fs.FlightRepo.Exists(flight) {
		return errors.FlightAlreadyExists()
	}

	if err := flight.IsValid(); err != nil {
		return err
	}

	if err := fs.FlightRepo.Save(flight); err != nil {
		return err
	}
	
	fs.flights.AddFlight(flight)

	return nil
}

func (fs *flightService) FindBestFlight(flight domain.Flight) error {
	flight.Cost = 1
	
	if err := flight.IsValid(); err != nil {
		return err
	}

	return fs.bookingService.FindBestFlight(flight.From, flight.To)
}
