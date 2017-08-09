package PocketMediaLimiter

import (
	"testing"
	"math"
	"time"
)

func Test_frequency_period(t *testing.T) {
	testSets := []struct{
		frequency
		period time.Duration
	}{
		{
			frequency:1.0,
			period:time.Second,
		},
		{
			frequency:2.0,
			period:time.Millisecond * 500,
		},
		{
			frequency:10.0,
			period:time.Millisecond * 100,
		},
		{
			frequency:0,
			period:time.Duration(math.Inf(1)),
		},
	}
	for _, testSet := range testSets {
		period := testSet.frequency.period()
		if testSet.period != period {
			t.Errorf("For a frequency of %f, expected period of %s but got %s", testSet.frequency, testSet.period, period)
		}
	}
}