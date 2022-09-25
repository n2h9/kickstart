package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	fabrics []*fabric

	expected int
}

var data []testData = []testData{
	{
		fabrics: []*fabric{
			{
				color:      "blue",
				durability: 2,
				id:         1,
			},
			{
				color:      "yellow",
				durability: 1,
				id:         2,
			},
		},
		expected: 0,
	},
	{
		fabrics: []*fabric{
			{
				color:      "blue",
				durability: 2,
				id:         1,
			},
			{
				color:      "brown",
				durability: 2,
				id:         2,
			},
		},
		expected: 2,
	},
	{
		fabrics: []*fabric{
			{
				color:      "red",
				durability: 2,
				id:         1,
			},
		},
		expected: 1,
	},
	{
		fabrics: []*fabric{
			{
				color:      "blue",
				durability: 1,
				id:         2,
			},
			{
				color:      "green",
				durability: 1,
				id:         4,
			},
			{
				color:      "orange",
				durability: 2,
				id:         5,
			},
			{
				color:      "red",
				durability: 3,
				id:         6,
			},
			{
				color:      "yellow",
				durability: 3,
				id:         7,
			},
		},
		expected: 5,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, numOfFabricks(item.fabrics))
	}
}
