package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const EPSILON = 0.000001

func TestSol(t *testing.T) {
	rates := []int{2000, 1500, 1900}
	res := highestPossibleRating(rates)
	assert.Equal(t, []int{1900, 2000, 2000}, res)
}
