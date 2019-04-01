package main

import (
	"net/http"
	"fmt"
	"os"
	"github.com/labstack/echo"
	"github.com/hdiomede/travel-scanner/domain"
	"github.com/hdiomede/travel-scanner/infrastructure/persistence"
)

var teste = persistence.RouteRepository{}

func main() {
	csvFile, _ := os.Open("/file.csv")
	teste.ReadFile(csvFile)
	e := echo.New()

	e.GET("/", hello)
	e.POST("/", newRoute)
	e.GET("/health", health)

	e.Logger.Fatal(e.Start(":8080"))

}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, domain.Route{"GRU", "MIA", 20})
}


func health(c echo.Context) error {

	rotas, _ := teste.All()

	fmt.Println(rotas[0])

	return c.String(http.StatusOK, "OK")
}


func newRoute(c echo.Context) (err error) {
	r := new(domain.Route)
	if err := c.Bind(r); err != nil {
		return err
	}

	teste.Save(r)
	rotas, _ := teste.All()

	return c.JSON(http.StatusOK, rotas)
}