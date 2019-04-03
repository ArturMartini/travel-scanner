package errors

import "fmt"

type AirportDoesNotExistsError struct {
	msg string
}

func (error *AirportDoesNotExistsError) Error() string {
	return error.msg
}

func AirportDoesNotExists(airportCode string) error {
	return &AirportDoesNotExistsError{fmt.Sprintf("Airport %s does not exists", airportCode)}
}

type FlightAlreadyExistsError struct {
	msg string
}

func (error *FlightAlreadyExistsError) Error() string {
	return error.msg
}

func FlightAlreadyExists() error {
	return &FlightAlreadyExistsError{"This flight already exists"}
}

type NoFlightFoundError struct {
	msg string
}

func (error *NoFlightFoundError) Error() string {
	return error.msg
}

func NoFlightFound() error {
	 return &NoFlightFoundError{"No flight found"}
}

type InvalidAirportCodeFormatError struct {
	msg string
}

func (error *InvalidAirportCodeFormatError) Error() string {
	return error.msg
}

func InvalidAirportCodeFormat() {
	return &InvalidAirportCodeFormatError{""}
}

type InvalidFlightCostError struct {
	msg string
}

func (error *InvalidFlightCostError) Error() string {
	return error.msg
}

func InvalidFlightCost() {
	return &InvalidFlightCostError{""}
}

