package domain


type Route struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Price int    `json:"price"`
}

type RouteRepository interface {
	Save(route *Route) error
	Exists(route *Route) bool
	All() ([]Route, error)
}