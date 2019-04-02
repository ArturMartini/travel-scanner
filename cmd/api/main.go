package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/hdiomede/travel-scanner/domain"
	"github.com/hdiomede/travel-scanner/application"
	"github.com/hdiomede/travel-scanner/infrastructure/persistence"
)

var teste = persistence.NewRouteRepository("/file.csv")
var service = application.NewRouteService(teste)

func main() {
	e := echo.New()

	e.GET("/health", health)
	e.GET("/routes", listRoutes)
	e.POST("/routes", newRoute)

	service.FindBestRoute()

	e.Logger.Fatal(e.Start(":8080"))
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func listRoutes(c echo.Context) error {
	routes, _ := service.All()
	return c.JSON(http.StatusOK, routes)
}

func newRoute(c echo.Context) (err error) {
	r := new(domain.Route)
	if err := c.Bind(r); err != nil {
		return err
	}

	service.SaveRoute(r)
	rotas, _ := service.All()

	return c.JSON(http.StatusOK, rotas)
}