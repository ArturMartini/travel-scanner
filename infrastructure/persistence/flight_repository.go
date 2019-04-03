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

type flightRepository struct {
	Flights []domain.Flight
	filename string
}

func NewFlightRepository(filename string) *flightRepository {
	fr := flightRepository{filename: filename}
	
	ok := fr.readFile()

	if ok != nil {
		log.Fatal("Error parsing csv file!")
		os.Exit(0)
	}

	return &fr
}

func (fr *flightRepository) readFile() error {
	csvFile, err := os.Open(fr.filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File %s will be created\n", fr.filename)
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

		var flight = domain.Flight{line[0], line[1], cost}

		if err := flight.IsValid(); err != nil {
			log.Fatal("Flight fields are invalid")
			return err
		}

		fr.Flights = append(fr.Flights, flight)
	}

	return nil
}

func (fr *flightRepository) Save(flight *domain.Flight) error {
	fr.Flights = append(fr.Flights, *flight)

	f, err := os.OpenFile(fr.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

func (fr *flightRepository) Exists(flight *domain.Flight) bool {
	for _, n := range fr.Flights {
		if flight.From == n.From && flight.To == n.To {
			return true
		}
	}

	return false
}

func (fr *flightRepository) All() ([]domain.Flight, error) {
	return fr.Flights, nil
}
