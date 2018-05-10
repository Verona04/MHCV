package main

import (
	"testing"
	"math"
)

func TestHsin(t *testing.T) {
	expected := math.Pow(math.Sin(5/2),2)
	actual := math.Pow(math.Sin(5/2),2)
	if actual != expected {
		t.Errorf("Test failed, expected: '%s' got: '%s'", expected, actual)
	}
}
