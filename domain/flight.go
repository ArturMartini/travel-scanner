package domain

import "github.com/hdiomede/travel-scanner/errors"

type Flight struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Cost int    `json:"cost"`
}

type Flights struct {
	Map  map[string]map[string]int
}

type FlightRepository interface {
	Save(flight *Flight) error
	Exists(flight *Flight) bool
	All() ([]Flight, error)
}

func (flight *Flight) IsValid() error {
	if err := validateAirportCode(flight.From); err != nil {
		return err
	}
	
	if err := validateAirportCode(flight.To); err != nil {
		return err
	}

	if err := validateFlightCost(flight.Cost); err != nil {
		return err
	}

	return nil
}

func validateAirportCode(airportCode string) error {
	if len(airportCode) != 3 {
		return errors.InvalidAirportCodeFormat()
	}

	return nil
}

func validateFlightCost(cost int) error {
	if cost <= 0 {
		return errors.InvalidFlightCost()
	}

	return nil
}


func (flights *Flights) AddFlight(flight *Flight) {
	child, ok := flights.Map[flight.From]
	if !ok {
		child = map[string]int{}
		flights.Map[flight.From] = child
	}

	child[flight.To] = flight.Cost
}