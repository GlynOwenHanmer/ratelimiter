package PocketMediaLimiter

import (
	"testing"
	"math"
)

func Test_frequency_period(t *testing.T) {
	testSets := []struct{
		frequency
		period float64
	}{
		{
			frequency:1.0,
			period:1.0,
		},
		{
			frequency:2.0,
			period:0.5,
		},
		{
			frequency:10.0,
			period:0.1,
		},
		{
			frequency:0,
			period:math.Inf(1),
		},
	}
	for _, testSet := range testSets {
		period := testSet.frequency.period()
		if testSet.period != period {
			t.Errorf("Expected period of %f but got %f", testSet.period, period)
		}
	}
}