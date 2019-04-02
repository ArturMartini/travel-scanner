package application


import (
	"fmt"
	"github.com/hdiomede/travel-scanner/domain"
)

type BookingService struct {
	Flights *domain.Flights
}

type vertex struct {
	Name string
	Perm bool
	Dist int
	Path string
}


func (bookingService *BookingService) FindBestFlight(origin string, target string) {
	nodes := make([]string, 0)
	for key := range bookingService.Flights.Map {
		nodes = append(nodes, key)
		for dest := range bookingService.Flights.Map[key] {
			nodes = append(nodes, dest)
		}
	}

	nodes = bookingService.uniqueElements(nodes)

	vertexList := make(map[string]*vertex, 0)

	var current string

	for _, node := range nodes {
		var temp = vertex{Name: node, Perm: false, Dist: 50000, Path: "-"}

		if node == origin {
			temp.Dist = 0
			current = node
		}

		vertexList[node] = &temp
	}
	
	for current != "" {
		vertexList[current].Perm = true
		var nextElements = bookingService.findNodes(vertexList[current].Name)
		
		for _, element := range nextElements {
			if (vertexList[element].Dist > vertexList[current].Dist + bookingService.Flights.Map[current][element]) {
				vertexList[element].Dist = vertexList[current].Dist + bookingService.Flights.Map[current][element]
				vertexList[element].Path = current
			}
		}
		current = bookingService.findNextCurrent(vertexList)
	}

	for _, vertex := range vertexList {
		fmt.Printf("[%s %t %d %s]\n",vertex.Name, vertex.Perm, vertex.Dist, vertex.Path)
	}

	bookingService.printPath(&vertexList, origin, target)	
}

func (bookingService *BookingService) printPath(vertexList *map[string]*vertex, origin string, dest string) {
	var route []string
	var current = dest
	

	for current != origin {
		route = append([]string{current}, route...)
		current = (*vertexList)[current].Path
	}

	route = append([]string{origin}, route...)

	fmt.Printf("Cheapest Route Price: %d\n", (*vertexList)[dest].Dist)
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

func (bookingService *BookingService) findNextCurrent(vertexList map[string]*vertex) string {
	var currentDist = 50001
	var currentNode = ""

	for _, vertex := range vertexList {
		if (!vertex.Perm && vertex.Dist < currentDist) {
			currentDist = vertex.Dist
			currentNode = vertex.Name
		}
	}
		
	return currentNode
}

func (bookingService *BookingService) findNodes(origin string) []string {
	candidates := make([]string, 0)

	for node := range bookingService.Flights.Map[origin] {
		candidates = append(candidates, node)
	}

	return candidates
}