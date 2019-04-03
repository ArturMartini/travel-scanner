package testing


import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/hdiomede/travel-scanner/domain"
	"github.com/hdiomede/travel-scanner/errors"
	"github.com/hdiomede/travel-scanner/application"
)

func TestFindBestFlightWhenAirportDoesNotExists(t *testing.T) {
	var flightMap = domain.Flights{make(map[string]map[string]int)}

	bookingService := &application.BookingService{&flightMap}

	err := bookingService.FindBestFlight(domain.Flight{From: "GRU", To: "MIA"})

	if assert.Error(t, err) {
		assert.Equal(t, errors.AirportDoesNotExists("GRU"), err)
	}
}


func TestFindBestFlightWhenNoFlightIsFound(t *testing.T) {
	var flightMap = domain.Flights{make(map[string]map[string]int)}
	
	flightMap.AddFlight(&domain.Flight{"GRU", "LIS", 20})
	flightMap.AddFlight(&domain.Flight{"CGH", "MIA", 10})

	bookingService := &application.BookingService{&flightMap}

	err := bookingService.FindBestFlight(domain.Flight{From: "GRU", To: "MIA"})

	if assert.Error(t, err) {
		assert.Equal(t, errors.NoFlightFound(), err)
	}
}

func TestFindBestFlight(t *testing.T) {
	var flightMap = domain.Flights{make(map[string]map[string]int)}
	
	flightMap.AddFlight(&domain.Flight{"GRU", "LIS", 20})
	flightMap.AddFlight(&domain.Flight{"LIS", "MIA", 10})

	bookingService := &application.BookingService{&flightMap}

	err := bookingService.FindBestFlight(domain.Flight{From: "GRU", To: "MIA"})

	assert.NoError(t, err, "Should not throw exception if conditions are met")
}
