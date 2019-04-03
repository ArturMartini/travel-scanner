package domain

type Flight struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Cost int    `json:"cost"`
}

type Flights struct {
	Map  map[string]map[string]int
}

type FlightRepository interface {
	Save(flight *Flight) error
	Exists(flight *Flight) bool
	All() ([]Flight, error)
}

func (flight *Flight) IsValid() bool {
	return flight.From != "" && flight.To != "" && flight.Cost > 0
}

func (flights *Flights) AddFlight(flight *Flight) {
	child, ok := flights.Map[flight.From]
	if !ok {
		child = map[string]int{}
		flights.Map[flight.From] = child
	}

	child[flight.To] = flight.Cost
}