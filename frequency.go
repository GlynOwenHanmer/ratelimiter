package ratelimiter

import (
	"time"
	"math"
	"errors"
)

// Frequency represents the rate at which something repeats in the units of Hertz, the number-of-events per second.
type Frequency float64

// NewFrequency returns a Frequency calculated from a given number of occurrences of an event over a given duration, also returning any errors that may be present in the given values.
// NewFrequency will return an error if the given Duration is smaller than a nanosecond.
func NewFrequency(occurrences uint64, duration time.Duration) (Frequency, error) {
	if duration < time.Nanosecond {
		return Frequency(math.MaxFloat64), errors.New("duration must be greater than 0")
	}
	return Frequency(float64(occurrences) / duration.Seconds()), nil
}

// period returns the period Duration for one cycle at a given Frequency.
func (f Frequency) period() time.Duration  {
	return time.Duration(1000000000.0 / float64(f))
}
