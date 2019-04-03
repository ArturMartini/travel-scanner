package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"net/http"
	"github.com/labstack/echo"
	"github.com/hdiomede/travel-scanner/domain"
	"github.com/hdiomede/travel-scanner/errors"
	"github.com/hdiomede/travel-scanner/application"
	"github.com/hdiomede/travel-scanner/infrastructure/persistence"
)

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

	e := echo.New()

	e.GET("/health", health)
	e.POST("/flights", newFlight)
	e.POST("/flights/search", searchFlight)

	e.Logger.Fatal(e.Start(":8080"))
}


func readInput() {
	reader := bufio.NewReader(os.Stdin)
  	fmt.Println("Please enter the route: ")
  	fmt.Println("---------------------")

  	for {
    	fmt.Print("Please enter the route: ")
    	text, _ := reader.ReadString('\n')
    	// convert CRLF to LF
    	text = strings.Replace(text, "\n", "", -1)

    	if strings.Compare("hi", text) == 0 {
      	fmt.Println("hello, Yourself")
    	}
  	}
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func newFlight(c echo.Context) (err error) {
	r := new(domain.Flight)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message" : "Invalid payload"})
	}

	errSave := service.SaveFlight(r)
	if errSave != nil {
		var status int

		switch t := errSave.(type) {
		default: 
			fmt.Println("Error")
			status = http.StatusInternalServerError
		case *errors.FlightAlreadyExistsError:
			fmt.Println("FlightAlreadyExists", t)
			status = http.StatusBadRequest
		case *errors.InvalidAirportCodeFormatError:
			fmt.Println("InvalidAirportCodeFormat", t)
			status = http.StatusBadRequest
		case *errors.InvalidFlightCostError:
			fmt.Println("InvalidFlightCost", t)
			status = http.StatusBadRequest
		case *errors.SaveFlightOperationError:
			fmt.Println("SaveFlightOperation", t)
			status = http.StatusInternalServerError
		}

		return c.JSON(status, map[string]string{"message" : errSave.Error()})
	}

	rotas, _ := service.All()

	return c.JSON(http.StatusOK, rotas)
}

func searchFlight(c echo.Context) (err error) {
	r := new(domain.Flight)
	if err := c.Bind(r); err != nil {
		fmt.Println("Invalid payload")
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message" : "Invalid payload"})
	}

	if errSearch := service.FindBestFlight(*r); errSearch != nil {
		var status int

		switch t := errSearch.(type) {
		default: 
			fmt.Println("Error")
			status = http.StatusInternalServerError
		case *errors.InvalidAirportCodeFormatError:
			fmt.Println("InvalidAirportCodeFormat", t)
			status = http.StatusBadRequest
		case *errors.NoFlightFoundError:
			fmt.Println("NoFlightFoundError", t)
			status = http.StatusNotFound
		case *errors.AirportDoesNotExistsError:
			fmt.Println("AirportDoesNotExistsError", t)
			status = http.StatusNotFound
		}

		return c.JSON(status, map[string]string{"message" : errSearch.Error()})
	}

	return c.String(http.StatusOK, "OK")
}