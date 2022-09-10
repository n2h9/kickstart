package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol(t *testing.T) {
	res := findMatchingPalindrome("abba")
	assert.Equal(t, "abba", res)
}

func TestSol1(t *testing.T) {
	res := findMatchingPalindrome("cccc")
	assert.Equal(t, "c", res)
}

func TestSol2(t *testing.T) {
	res := findMatchingPalindrome("cdccdc")
	assert.Equal(t, "cdc", res)
}

func TestSol3(t *testing.T) {
	res := findMatchingPalindrome("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	assert.Equal(t, "a", res)
}
