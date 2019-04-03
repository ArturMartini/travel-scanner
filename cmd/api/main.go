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

var teste = persistence.NewFlightRepository("/app/file.csv")
var service = application.NewFlightService(teste)

func main() {
	e := echo.New()

	e.GET("/health", health)
	e.GET("/flights", listFlights)
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

func listFlights(c echo.Context) error {
	routes, _ := service.All()
	return c.JSON(http.StatusOK, routes)
}

func newFlight(c echo.Context) (err error) {
	r := new(domain.Flight)
	if err := c.Bind(r); err != nil {
		return err
	}

	errSave := service.SaveFlight(r)
	if errSave != nil {
		switch t := errSave.(type) {
		default: 
			fmt.Println("Error")
		case *errors.FlightAlreadyExistsError:
			fmt.Println("FlightAlreadyExists", t)
		case *errors.InvalidAirportCodeFormatError:
			fmt.Println("InvalidAirportCodeFormat", t)
		case *errors.InvalidFlightCostError:
			fmt.Println("InvalidFlightCost", t)
		case *errors.SaveFlightOperationError:
			fmt.Println("SaveFlightOperation", t)
		}
	}

	rotas, _ := service.All()

	return c.JSON(http.StatusOK, rotas)
}

func searchFlight(c echo.Context) (err error) {
	r := new(domain.Flight)
	if err := c.Bind(r); err != nil {
		fmt.Println("Invalid payload")
		return err
	}

	if errSearch := service.FindBestFlight(*r); errSearch != nil {
		switch t := errSearch.(type) {
		default: 
			fmt.Println("Error")
		case *errors.InvalidAirportCodeFormatError:
			fmt.Println("InvalidAirportCodeFormat", t)
		case *errors.NoFlightFoundError:
			fmt.Println("NoFlightFoundError", t)
		case *errors.AirportDoesNotExistsError:
			fmt.Println("AirportDoesNotExistsError", t)
		}
	}

	return c.String(http.StatusOK, "OK")
}