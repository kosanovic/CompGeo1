package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	// Simples Beispiel analog zum Skript
	testGraphs := []Graph{
		{
			p1X: -6,
			p1Y: 1,
			p2X: 5,
			p2Y: 0,
		},
		{
			p1X: -7,
			p1Y: -1,
			p2X: 3,
			p2Y: 6,
		},
		{
			p1X: -1,
			p1Y: 6,
			p2X: 4,
			p2Y: -2,
		},
	}
	result := amountOfInterceptingGraphs(testGraphs)

	expected := 3
	if result != expected {
		t.Errorf("Amount of intercepting graphs returned %d intersections, expected %d", result, expected)
	}
}
