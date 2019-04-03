package persistence

import (
	"io"
	"os"
	"fmt"
	"log"
	"strconv"
    "encoding/csv"
	"github.com/hdiomede/travel-scanner/errors"
	"github.com/hdiomede/travel-scanner/domain"
)

type FlightRepository struct {
	Flights []domain.Flight
	filename string
}

func NewFlightRepository(filename string) *FlightRepository {
	flightRepository := FlightRepository{filename: filename}
	
	ok := flightRepository.readFile()

	if ok != nil {
		log.Fatal("Error parsing csv file!")
		os.Exit(0)
	}

	return &flightRepository
}

func (flightRepository *FlightRepository) readFile() error {
	csvFile, err := os.Open(flightRepository.filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File %s will be created\n", flightRepository.filename)
			return nil
		}

		fmt.Println(err)
		return err
	}

	reader := csv.NewReader(csvFile)

	for {
		line, err := reader.Read()
		if err != nil{
			if err == io.EOF {
				break
			}
			
			return errors.CsvParse()
		}
		cost, _ := strconv.Atoi(line[2])

		flightRepository.Flights = append(flightRepository.Flights, domain.Flight{line[0], line[1], cost})
	}

	return nil
}

func (flightRepository *FlightRepository) Save(flight *domain.Flight) error {
	flightRepository.Flights = append(flightRepository.Flights, *flight)

	f, err := os.OpenFile(flightRepository.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return errors.SaveFlightOperation()
    }
	if _, err := f.Write([]byte(fmt.Sprintf("\n%s,%s,%d", flight.From, flight.To, flight.Cost))); err != nil {
        log.Fatal(err)
		return errors.SaveFlightOperation()
    }
    if err := f.Close(); err != nil {
        log.Fatal(err)
		return errors.SaveFlightOperation()
    }

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
