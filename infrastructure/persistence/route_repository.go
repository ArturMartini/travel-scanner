package persistence

import (
	"io"
	"os"
	"fmt"
    "encoding/csv"
	"github.com/hdiomede/travel-scanner/domain"
)

type RouteRepository struct {
	Routes []domain.Route
}


func (r *RouteRepository) ReadFile(csvFile *os.File) error {
	reader := csv.NewReader(csvFile)
	//var routes []domain.Route

	for {
		line, err := reader.Read()
		if err == io.EOF {
            break
        }
		fmt.Println(line)
	}

	return nil
}

func (r *RouteRepository) Save(route *domain.Route) error {
	r.Routes = append(r.Routes, *route)

	return nil
}

func (r *RouteRepository) All() ([]domain.Route, error) {
	return r.Routes, nil
}
