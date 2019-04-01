package persistence

import (
	"io"
	"os"
	"fmt"
	"strconv"
    "encoding/csv"
	"github.com/hdiomede/travel-scanner/domain"
)

type RouteRepository struct {
	Routes []domain.Route
}


func (r *RouteRepository) ReadFile(filename string) error {
	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(csvFile)

	for {
		line, err := reader.Read()
		if err == io.EOF {
            break
        }
		fmt.Println(line)
		price, _ := strconv.Atoi(line[2])

		r.Routes = append(r.Routes, domain.Route{line[0], line[1], price})
	}

	return nil
}

func (r *RouteRepository) Save(route *domain.Route) error {
	r.Routes = append(r.Routes, *route)

	return nil
}

func (r *RouteRepository) Exists(route *domain.Route) bool {
	for _, n := range r.Routes {
		if route.From == n.From && route.To == n.To {
			return true
		}
	}

	return false
}

func (r *RouteRepository) All() ([]domain.Route, error) {
	return r.Routes, nil
}
