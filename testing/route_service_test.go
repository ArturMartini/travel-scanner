package testing

/*
import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/hdiomede/travel-scanner/domain"
	"github.com/hdiomede/travel-scanner/application"
)

type routeRepositoryMock struct {
	mock.Mock
}

func (m *routeRepositoryMock) All() ([]domain.Route, error) {
	args := m.Called()
	return args.Get(0).([]domain.Route), nil
}

func (m *routeRepositoryMock) Exists(route *domain.Route) bool {
	args := m.Called()

	return args.Bool(0)
}

func (m *routeRepositoryMock) Save(route *domain.Route) error {
	return nil
}

func TestLoadRoutes(t *testing.T) {
	var routesList []domain.Route
	routesList = append(routesList, domain.Route{"MIA", "GRU", 10})

	routeRepoMock := new(routeRepositoryMock)
	routeRepoMock.On("All").Return(routesList, nil)
	var service = application.NewRouteService(routeRepoMock)

	routes, _ := service.All()
	assert.Equal(t, routes[0].From, "MIA", "From should be MIA")
}

func TestSaveWhenRouteExists(t *testing.T) {
	routeRepoMock := new(routeRepositoryMock)
	routeRepoMock.On("Exists").Return(true)
	//var service = application.NewRouteService(routeRepoMock)

	assert.Equal(t, routeRepoMock.Exists(nil), true, "Should be true")
}
*/