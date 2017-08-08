package PocketMediaLimiter

import (
	"errors"
	"time"
	"math"
)

// Limiter is a token-bucket implemented rate limiter.
// A Limiter is configured to have an average rate at which the tokens can be retrieved from the bucket and a maximum burst rate at which the tokens can be taken out of the bucket.
// The maximum burst rate of the represents the depth of the token bucket and is the number of times that a token can be taken from the bucket if the bucket were to be taken from completely full to completely empty in quick succession, without time being given for the bucket tokens to be replenished in between any tokens being retrieved.
type Limiter struct {
	rate frequency
	tokens uint64
	burst uint64
}

// Rate returns the average rate, measured as events per second, at which the limiter is to allow an event to be triggered.
func (limiter Limiter) Rate() float64 {
	return float64(limiter.rate)
}

// NewLimiter returns a Limiter with the given rate and an error if there is a problem with the rate provided.
// NewLimiter will return a Limiter with a rate of 0 and an error if a negative rate is given.
func NewLimiter(rate float64, burst uint64) (*Limiter, error) {
	if rate < 0 {
		return &Limiter{}, errors.New("Limiter rate cannot be negative.")
	}
	frequency := frequency(rate)
	limiter := Limiter{rate:frequency, tokens:burst, burst:burst}
	createIncrementer(frequency, &limiter)
	return &limiter, nil
}

// createIncrementer starts a goroutine that will increment the Limiter token bucket at a rate equal to the rate of the Limiter
func createIncrementer(frequency frequency, limiter *Limiter) {
	ticker := createTicker(frequency)
	go func() {
		for range ticker.C {
			limiter.increment()
		}
	}()
}

// createTicker returns a ticker that will tick at a given frequency
func createTicker(frequency frequency) *time.Ticker {
	var duration time.Duration
	if frequency <= 0.0 {
		duration = time.Duration(math.MaxInt64)
	} else {
		duration = time.Duration(frequency.period() * float64(time.Second.Nanoseconds()))
	}
	return time.NewTicker(duration)
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

// increment increments the number of tokens available in the bucket if there is space in the bucket to do so.
func (limiter *Limiter) increment() {
	if limiter.tokens >= limiter.burst {
		return
	}
	limiter.tokens++
}

// decrement decrements the number of tokens available in the bucket if there are any left.
func (limiter *Limiter) decrement() {
	if limiter.tokens == 0 {
		return
	}
	limiter.tokens--
}

// hasTokens returns true if the limiter has available tokens in its bucket
func (limiter *Limiter) hasTokens() bool {
	return limiter.tokens > 0
}