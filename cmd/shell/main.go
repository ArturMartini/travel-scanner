package main

import (
	"fmt"
	"os"
	"github.com/hdiomede/travel-scanner/domain"
	"github.com/hdiomede/travel-scanner/errors"
	"github.com/hdiomede/travel-scanner/application"
	"github.com/hdiomede/travel-scanner/infrastructure/persistence"
)

/*
type BookingResult struct {
	Cost int `json:"cost"`
	Route string `json:"route"`
}
*/

var repo domain.FlightRepository
var service application.FlightService

func main() {
    argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 0 {
		fmt.Println("No filename specified")
		os.Exit(0)
	}

	repo = persistence.NewFlightRepository(argsWithoutProg[0])
	service = application.NewFlightService(repo)

	fmt.Println(argsWithoutProg)

	for {
		var origin string
		var dest string
		
		fmt.Print("Please enter route: ")
		fmt.Scanf("%s %s\n", &origin, &dest)
		fmt.Println("Origin: %s\n", origin)
		fmt.Println("Dest: %s\n", dest)
		fmt.Printf("%s\n", searchFlight(origin, dest))
  	}
}

func searchFlight(origin string, dest string) string {
	r := &domain.Flight{From: origin, To: dest}

	path, cost, errSearch := service.FindBestFlight(*r)

	if  errSearch != nil {
		var message string

		switch errSearch.(type) {
		default: 
			message = "Error"
		case *errors.InvalidAirportCodeFormatError:
			message = errSearch.Error()
		case *errors.NoFlightFoundError:
			message = errSearch.Error()
		case *errors.AirportDoesNotExistsError:
			message = errSearch.Error()
		}

		return message
	}

	return fmt.Sprintf("%s > %d", path, cost)
}