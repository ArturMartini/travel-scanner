package domain

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRoute(t *testing.T) {
	route := &Route{"GRU", "MIA", 20}
	assert.Equal(t, route.From, "GRU", "From should be GRU")
	assert.Equal(t, route.To, "MIA", "To should be MIA")
	assert.Equal(t, route.Price, 20, "Price should be 20")
}