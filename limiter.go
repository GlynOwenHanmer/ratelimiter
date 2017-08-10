package ratelimiter

import (
	"errors"
)

// Limiter is a token-bucket implemented rate limiter.
// A Limiter is configured to have an average rate at which the tokens can be retrieved from the bucket and a maximum burst rate at which the tokens can be taken out of the bucket.
// The maximum burst rate of the represents the depth of the token bucket and is the number of times that a token can be taken from the bucket if the bucket were to be taken from completely full to completely empty in quick succession, without time being given for the bucket tokens to be replenished in between any tokens being retrieved.
type Limiter struct {
	rate               Frequency
	tokenBucket
	tickingIncrementer TickingIncrementer
}

// Rate returns the average rate, measured as events per second, at which the limiter is to allow an event to be triggered.
func (limiter Limiter) Rate() Frequency {
	return limiter.rate
}

// NewLimiter returns a Limiter with the given rate and an error if there is a problem with the rate provided.
// NewLimiter will return a Limiter with a rate of 0 and an error if a negative rate is given.
func NewLimiter(rate Frequency, burst uint64) (*Limiter, error) {
	if rate < 0.0 {
		return &Limiter{}, errors.New("Limiter rate cannot be negative.")
	}
	frequency := Frequency(rate)
	bucket := tokenBucket{tokens:1, depth:burst}
	limiter := Limiter{rate:frequency, tokenBucket:bucket}

	var ti TickingIncrementer
	var err error
	if frequency <= 0.0 {
		ti = TickingIncrementer{}
	} else {
		ti, err = NewTickingIncrementer(&limiter.tokenBucket, frequency.period())
		if err != nil {
			return &Limiter{}, err
		}
	}
	limiter.tickingIncrementer = ti
	return &limiter, nil
}

// Allow returns true if the event can be invoked within the allowed rate of the Limiter.
// Allow will always return false for a Limiter with a rate of 0.
func (limiter *Limiter) Allow() bool {
	if limiter.rate <= 0.0 {
		return false
	}
	if !limiter.hasTokens() {
		return false
	}
	limiter.decrement()
	return true
}
