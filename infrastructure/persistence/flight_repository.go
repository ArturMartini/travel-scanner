package persistence

import (
	"io"
	"os"
	"fmt"
	"strconv"
    "encoding/csv"
	"github.com/hdiomede/travel-scanner/domain"
)

type FlightRepository struct {
	Flights []domain.Flight
}

func NewFlightRepository(filename string) *FlightRepository {
	flightRepository := FlightRepository{}
	
	flightRepository.readFile(filename)

	return &flightRepository
}

func (flightRepository *FlightRepository) readFile(filename string) error {
	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(csvFile)

	for {
		line, err := reader.Read()
		if err == io.EOF {
            break
        }
		fmt.Println(line)
		cost, _ := strconv.Atoi(line[2])

		flightRepository.Flights = append(flightRepository.Flights, domain.Flight{line[0], line[1], cost})
	}

	return nil
}

func (flightRepository *FlightRepository) Save(flight *domain.Flight) error {
	flightRepository.Flights = append(flightRepository.Flights, *flight)

	return nil
}

func (flightRepository *FlightRepository) Exists(flight *domain.Flight) bool {
	for _, n := range flightRepository.Flights {
		if flight.From == n.From && flight.To == n.To {
			return true
		}
	}

	return false
}

func (flightRepository *FlightRepository) All() ([]domain.Flight, error) {
	return flightRepository.Flights, nil
}
