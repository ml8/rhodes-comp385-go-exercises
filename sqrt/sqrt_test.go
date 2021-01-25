package main

import (
	"math"
	"testing"
)

const closeEnough = 0.001

func TestSqrt(t *testing.T) {
	for i := 1.0; i < 10_000.0; i += 10 {
		actual, err := sqrt(i)
		if err != nil {
			t.Errorf("Unexpected error for %v: %v", i, err)
		}
		expected := math.Sqrt(i)
		if math.Abs(actual-expected) > closeEnough {
			t.Errorf("%v too far from reference value %v", actual, expected)
		}
	}
}

func TestNegativeSqrt(t *testing.T) {
	v, err := sqrt(-2)
	if err == nil {
		t.Errorf("Expected error for sqrt(-2)! Actual: %v", v)
	}
}
