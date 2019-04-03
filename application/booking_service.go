package application


import (
	"fmt"
	"github.com/hdiomede/travel-scanner/domain"
	"github.com/hdiomede/travel-scanner/errors"
)

type BookingService struct {
	Flights *domain.Flights
}

type vertex struct {
	Name string
	Perm bool
	Cost int
	Path string
}

const max_cost int = 500000000

func (bookingService *BookingService) FindBestFlight(origin string, target string) error {
	nodes := make([]string, 0)
	for key := range bookingService.Flights.Map {
		nodes = append(nodes, key)
		for dest := range bookingService.Flights.Map[key] {
			nodes = append(nodes, dest)
		}
	}

	nodes = bookingService.uniqueElements(nodes)

	for _, airportCode := range []string{origin, target} {
		if !bookingService.checkAirPortExists(nodes, airportCode) {
			return errors.AirportDoesNotExists(airportCode)	
		}
	}
	
	vertexList := make(map[string]*vertex, 0)

	var current string

	for _, node := range nodes {
		var temp = vertex{Name: node, Perm: false, Cost: max_cost, Path: "-"}

		if node == origin {
			temp.Cost = 0
			current = node
		}

		vertexList[node] = &temp
	}
	
	var candidates = true

	for candidates {
		vertexList[current].Perm = true
		var nextElements = bookingService.findNodes(vertexList[current].Name)
		
		for _, element := range nextElements {
			if (vertexList[element].Cost > vertexList[current].Cost + bookingService.Flights.Map[current][element]) {
				vertexList[element].Cost = vertexList[current].Cost + bookingService.Flights.Map[current][element]
				vertexList[element].Path = current
			}
		}
		current, candidates = bookingService.findNextCurrent(vertexList)
	}

	if bookingService.checkFlightNotFound(&vertexList, target) {
		return errors.NoFlightFound()
	}

	bookingService.printPath(&vertexList, origin, target)

	return nil
}

func (bookingService *BookingService) checkAirPortExists(airports []string, airport string) bool {
	for _, value := range airports {
		if value == airport {
			return true
		}
	}

	return false
}

func (bookingService *BookingService) checkFlightNotFound(vertexList *map[string]*vertex, target string) bool {
	return (*vertexList)[target].Cost == max_cost
}

func (bookingService *BookingService) printPath(vertexList *map[string]*vertex, origin string, dest string) {
	var route []string
	var current = dest
	

	for current != origin {
		route = append([]string{current}, route...)
		current = (*vertexList)[current].Path
	}

	route = append([]string{origin}, route...)

	fmt.Printf("Cheapest Route Price: %d\n", (*vertexList)[dest].Cost)
	fmt.Println(route)
}

func (bookingService *BookingService) uniqueElements(vector []string) []string {
	elements := make(map[string]bool)
	list := []string{}
	for _, entry := range vector {
		if _, value := elements[entry]; !value {
			elements[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

func (bookingService *BookingService) findNextCurrent(vertexList map[string]*vertex) (string, bool) {
	var anyCandidate = false
	var currentCost = max_cost
	var currentNode = ""

	for _, vertex := range vertexList {
		if (!vertex.Perm && vertex.Cost < currentCost) {
			currentCost = vertex.Cost
			currentNode = vertex.Name
			anyCandidate = true
		}
	}
		
	return currentNode, anyCandidate
}

func (bookingService *BookingService) findNodes(origin string) []string {
	candidates := make([]string, 0)

	for node := range bookingService.Flights.Map[origin] {
		candidates = append(candidates, node)
	}

	return candidates
}