package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	n           int
	connections [][]int
	queries     []int

	expected int
}

var data []testData = []testData{
	{
		n:           3,
		connections: [][]int{{2, 1}, {1, 3}},
		queries:     []int{1, 2},
		expected:    1,
	},
	{
		n:           4,
		connections: [][]int{{4, 2}, {2, 1}, {3, 1}},
		queries:     []int{3, 3, 3, 3},
		expected:    4,
	},
	{
		n:           5,
		connections: [][]int{{4, 2}, {2, 1}, {3, 1}, {5, 3}},
		queries:     []int{4, 5},
		expected:    1,
	},
	{
		n:           1,
		connections: [][]int{},
		queries:     []int{1},
		expected:    1,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, numOfComplitelyFilledContainers(item.n, item.connections, item.queries))
	}
}
