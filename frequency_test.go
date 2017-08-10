package ratelimiter_test

import (
	"testing"
	"time"
	"fmt"
	"math"
	"github.com/GlynOwenHanmer/ratelimiter"
)

func TestNewFrequency(t *testing.T) {
	testSets := []struct{
		occurences uint64
		time.Duration
		ratelimiter.Frequency
		expectError bool
	}{
		{
			occurences:0,
			Duration:1,
			Frequency:0.0,
			expectError:false,
		},
		{
			occurences:0,
			Duration:0,
			Frequency:ratelimiter.Frequency(math.MaxFloat64),
			expectError:true,
		},
		{
			occurences:0,
			Duration:-1,
			Frequency:ratelimiter.Frequency(math.MaxFloat64),
			expectError:true,
		},
		{
			occurences:10,
			Duration:time.Second,
			expectError:false,
			Frequency:ratelimiter.Frequency(10),
		},
		{
			occurences:10,
			Duration:time.Millisecond,
			expectError:false,
			Frequency:ratelimiter.Frequency(10000),
		},
	}
	for _, testSet := range testSets {
		freq, err := ratelimiter.NewFrequency(testSet.occurences, testSet.Duration)
		switch {
		case err == nil && testSet.expectError:
			t.Errorf("Expected error but received none.")
		case err != nil && !testSet.expectError:
			t.Errorf("Expected no error but received: %s", err.Error())
		}
		if freq != testSet.Frequency {
			testDetails := fmt.Sprintf("occurences: %d, duration %s", testSet.occurences, testSet.Duration)
			t.Errorf("Unexpected Frequency result.\n\tExpected: %f\n\tActual  : %f\n\ttest: %s", testSet.Frequency, freq, testDetails)
		}
	}
}

func ExampleNewFrequency() {
	rate, err := ratelimiter.NewFrequency(250, time.Second * 2)
	if err != nil {
		fmt.Printf("Unable to create Frequency: %s", err.Error())
		return
	}
	fmt.Printf("%.2f", rate)
	// Output: 125.00
}