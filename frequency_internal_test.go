package ratelimiter

import (
	"testing"
	"math"
	"time"
)

func Test_frequency_period(t *testing.T) {
	testSets := []struct{
		Frequency
		period time.Duration
	}{
		{
			Frequency: 1.0,
			period:    time.Second,
		},
		{
			Frequency: 2.0,
			period:    time.Millisecond * 500,
		},
		{
			Frequency: 10.0,
			period:    time.Millisecond * 100,
		},
		{
			Frequency: 0,
			period:    time.Duration(math.Inf(1)),
		},
	}
	for _, testSet := range testSets {
		period := testSet.Frequency.period()
		if testSet.period != period {
			t.Errorf("For a Frequency of %f, expected period of %s but got %s", testSet.Frequency, testSet.period, period)
		}
	}
}