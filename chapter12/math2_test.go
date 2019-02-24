package main

import (
	"golang-book/chapter11/math"
	"testing"
)

type testpair struct {
	values []float64
	average float64
}

var tests = []testpair{
	{ []float64{1,2}, 1.5 },
	{ []float64{1,1,1,1,1,1}, 1 },
	{ []float64{-1,1}, 0 },
	{ []float64{0,0}, 0 },
}

func TestAverage2(t *testing.T) {
	for _, pair := range tests {
		v := math.Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}
