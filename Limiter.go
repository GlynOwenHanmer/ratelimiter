package PocketMediaLimiter

import "errors"

type Limiter struct {
	rate float64
}

// Rate returns the average rate, measured as events per second, at which the limiter is to allow an event to be triggered.
func (limiter Limiter) Rate() float64 {
	return limiter.rate
}

func (limiter Limiter) Allow() bool {
	return false
}

// NewLimiter returns a Limiter with the given rate and an error if there is a problem with the rate provided.
// NewLimiter will return a Limiter with a rate of 0 and an error if a negative rate is given.
func NewLimiter(rate float64) (Limiter, error) {
	if rate < 0 {
		return Limiter{}, errors.New("Limiter rate cannot be negative.")
	}
	return Limiter{rate:rate}, nil
}