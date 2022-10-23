package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	inpt *input

	expected int
}

var data []testData = []testData{
	{
		inpt: &input{
			m: 2,
			n: 3,
			p: 1,
			scores: [][]int{
				{1000, 2000, 3000},
				{1500, 1500, 3000},
			},
		},
		expected: 500,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, minTotalAdditionalSteps(item.inpt))
	}
}
