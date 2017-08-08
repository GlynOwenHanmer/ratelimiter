package PocketMediaLimiter

import "errors"

type Limiter struct {
	rate float64
	tokens bool
}

// Rate returns the average rate, measured as events per second, at which the limiter is to allow an event to be triggered.
func (limiter Limiter) Rate() float64 {
	return limiter.rate
}

// NewLimiter returns a Limiter with the given rate and an error if there is a problem with the rate provided.
// NewLimiter will return a Limiter with a rate of 0 and an error if a negative rate is given.
func NewLimiter(rate float64) (Limiter, error) {
	if rate < 0 {
		return Limiter{}, errors.New("Limiter rate cannot be negative.")
	}
	return Limiter{rate:rate, tokens:true}, nil
}

// Allow returns true if the event can be invoked within the allowed rate of the Limiter.
// Allow will always return false for a Limiter with a rate of 0.
func (limiter *Limiter) Allow() bool {
	if limiter.rate <= 0.0 {
		return false
	}
	if !limiter.tokens {
		return false
	}
	limiter.decrement()
	return true
}

// increment increments the number of tokens available in the bucket if there is space in the bucket to do so.
func (limiter *Limiter) increment() {
	limiter.tokens = true
}

//decrement decrements the number of tokens available in the bucket if there are any left.
func (limiter *Limiter) decrement() {
	limiter.tokens = false
}