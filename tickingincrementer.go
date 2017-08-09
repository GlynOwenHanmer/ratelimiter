package PocketMediaLimiter

import (
	"time"
	"errors"
)

// TickingIncrementer is contains an Incrementer that will be incremented at regular intervals of a given time.Duration.
type TickingIncrementer struct {
	// Incrementer to be incremented
	incrementer Incrementer
	// Duration at which the incrementing happens
	duration time.Duration
	// Ticker that will tick to then do the incrementing.
	ticker *time.Ticker
}

// NewTickingIncrementer returns a TickingIncrementer that will increment the given Incrementer at regular intervals of the given time.Duration.
// If a nil Incrementer or a non-positive time.Duration is given, an empty TickingIncrementer will be returned along with an error describing the issue with the Incrementer or time.Duration.
func NewTickingIncrementer(incrementer Incrementer, duration time.Duration) (TickingIncrementer, error){
	if incrementer == nil {
		return TickingIncrementer{}, errors.New("No Incrementer provided.")
	}
	if duration < time.Nanosecond {
		return TickingIncrementer{}, errors.New("Duration must be positive.")
	}
	ticker := time.NewTicker(duration)
	go func() {
		for range ticker.C {
			incrementer.Increment()
		}
	}()
	return TickingIncrementer{incrementer:incrementer, duration:duration, ticker:ticker}, nil
}

// Stop is used to stop the TimingIncrementer from ticking and therefore stop the Incrementer from being further incremented.
func (ti TickingIncrementer) Stop() {
	if ti.ticker == nil {
		return
	}
	ti.ticker.Stop()
}