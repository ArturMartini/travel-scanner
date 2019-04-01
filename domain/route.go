package domain


type Route struct {
	From  string
	To    string
	Price int
}

type RouteRepository interface {
	Save(route *Route) error
	All() ([]Route, error)
}