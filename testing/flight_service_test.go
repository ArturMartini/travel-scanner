package testing

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/hdiomede/travel-scanner/errors"
	"github.com/hdiomede/travel-scanner/domain"
	"github.com/hdiomede/travel-scanner/application"
)

type flightRepositoryMock struct {
	mock.Mock
}

func (m *flightRepositoryMock) All() ([]domain.Flight, error) {
	args := m.Called()
	return args.Get(0).([]domain.Flight), nil
}

func (m *flightRepositoryMock) Exists(flight *domain.Flight) bool {
	args := m.Called()

	return args.Bool(0)
}

func (m *flightRepositoryMock) Save(flight *domain.Flight) error {
	return nil
}

func TestAll(t *testing.T) {
	flightsList := []domain.Flight{domain.Flight{"GRU", "LIS", 20}, domain.Flight{"LIS", "GRU", 10}}

	flightRepoMock := new(flightRepositoryMock)
	flightRepoMock.On("All").Return(flightsList, nil)

	var service = application.NewFlightService(flightRepoMock)

	flights, _ := service.All()
	assert.Equal(t, flights[0].From, "GRU", "From should be GRU")
	assert.Equal(t, flights[0].To, "LIS", "To should be LIS")
	assert.Equal(t, flights[0].Cost, 20, "Cost should be 20")

	assert.Equal(t, flights[1].From, "LIS", "From should be GRU")
	assert.Equal(t, flights[1].To, "GRU", "To should be LIS")
	assert.Equal(t, flights[1].Cost, 10, "Cost should be 20")
}

func TestSaveFlightAlreadyExists(t *testing.T) {
	flight := domain.Flight{"GRU", "LIS", 20}
	flightsList := []domain.Flight{domain.Flight{"GRU", "LIS", 20}, domain.Flight{"LIS", "GRU", 10}}

	flightRepoMock := new(flightRepositoryMock)
	flightRepoMock.On("All").Return(flightsList, nil)
	flightRepoMock.On("Exists").Return(true)

	var service = application.NewFlightService(flightRepoMock)

	err := service.SaveFlight(&flight)

	if assert.Error(t, err) {
		assert.Equal(t, errors.FlightAlreadyExists(), err)
	}
}

func TestSaveFlightInvalidAirportCodeFormat(t *testing.T) {
	flight := domain.Flight{"", "LIS", 20}
	flightsList := []domain.Flight{domain.Flight{"GRU", "LIS", 20}, domain.Flight{"LIS", "GRU", 10}}

	flightRepoMock := new(flightRepositoryMock)
	flightRepoMock.On("All").Return(flightsList, nil)
	flightRepoMock.On("Exists").Return(false)

	var service = application.NewFlightService(flightRepoMock)

	err := service.SaveFlight(&flight)

	if assert.Error(t, err) {
		assert.Equal(t, errors.InvalidAirportCodeFormat(), err)
	}
}

func TestSaveFlightInvalidFlightCost(t *testing.T) {
	flight := domain.Flight{"GRU", "LIS", -20}
	flightsList := []domain.Flight{domain.Flight{"GRU", "LIS", 20}, domain.Flight{"LIS", "GRU", 10}}

	flightRepoMock := new(flightRepositoryMock)
	flightRepoMock.On("All").Return(flightsList, nil)
	flightRepoMock.On("Exists").Return(false)

	var service = application.NewFlightService(flightRepoMock)

	err := service.SaveFlight(&flight)

	if assert.Error(t, err) {
		assert.Equal(t, errors.InvalidFlightCost(), err)
	}
}

func TestSaveFlight(t *testing.T) {

	flight := domain.Flight{"GRU", "LIS", 20}
	flightsList := []domain.Flight{domain.Flight{"GRU", "LIS", 20}, domain.Flight{"LIS", "GRU", 10}}

	flightRepoMock := new(flightRepositoryMock)
	flightRepoMock.On("All").Return(flightsList, nil)
	flightRepoMock.On("Exists").Return(false)
	flightRepoMock.On("Save").Return(nil)

	var service = application.NewFlightService(flightRepoMock)

	err := service.SaveFlight(&flight)

	assert.NoError(t, err, "Should not throw error")
}