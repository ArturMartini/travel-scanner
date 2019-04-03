package testing

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/hdiomede/travel-scanner/errors"
	"github.com/hdiomede/travel-scanner/domain"
)

func TestFlightStruct(t *testing.T) {
	flight := &domain.Flight{"GRU", "MIA", 20}

	assert.Equal(t, flight.From, "GRU", "From should be GRU")
	assert.Equal(t, flight.To, "MIA", "To should be MIA")
	assert.Equal(t, flight.Cost, 20, "Price should be 20")
}

func TestValidFlight(t *testing.T) {
	flight := &domain.Flight{"GRU", "MIA", 20}
	err := flight.IsValid()

	assert.NoError(t, err, "Should not return error if all fields are valid")
}

func TestFromInvalidAirportCodeFormat(t *testing.T) {
	flight := &domain.Flight{"", "MIA", 20}
	err := flight.IsValid()

	if assert.Error(t, err) {
		assert.Equal(t, errors.InvalidAirportCodeFormat(), err)
	}
}

func TestToInvalidAirportCodeFormat(t *testing.T) {
	flight := &domain.Flight{"GRU", "MI", 20}
	err := flight.IsValid()

	if assert.Error(t, err) {
		assert.Equal(t, errors.InvalidAirportCodeFormat(), err)
	}
}

func TestInvalidFlightCost(t *testing.T) {
	flight := &domain.Flight{"GRU", "MIA", -20}
	err := flight.IsValid()

	if assert.Error(t, err) {
		assert.Equal(t, errors.InvalidFlightCost(), err)
	}
}

func TestAddFlight(t *testing.T) {
	flight := &domain.Flight{"GRU", "MIA", 20}
	flights := &domain.Flights{make(map[string]map[string]int)}

	flights.AddFlight(flight)

	assert.Equal(t, flights.Map[flight.From][flight.To], flight.Cost)
}

func TestAddFlightWhenNested(t *testing.T) {
	flight := &domain.Flight{"GRU", "MIA", 20}
	otherFlight := &domain.Flight{"GRU", "LIS", 10}
	flights := &domain.Flights{make(map[string]map[string]int)}

	flights.AddFlight(flight)	
	flights.AddFlight(otherFlight)

	assert.Equal(t, flights.Map[flight.From][flight.To], flight.Cost)
	assert.Equal(t, flights.Map[otherFlight.From][otherFlight.To], otherFlight.Cost)
}

func TestAddFlightWhenNewFrom(t *testing.T) {
	flight := &domain.Flight{"GRU", "MIA", 20}
	otherFlight := &domain.Flight{"LIS", "GRU", 10}
	flights := &domain.Flights{make(map[string]map[string]int)}

	flights.AddFlight(flight)	
	flights.AddFlight(otherFlight)

	assert.Equal(t, flights.Map[flight.From][flight.To], flight.Cost)
	assert.Equal(t, flights.Map[otherFlight.From][otherFlight.To], otherFlight.Cost)
}