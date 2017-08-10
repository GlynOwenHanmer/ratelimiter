package ratelimiter

// Incrementer is an interface for something that can increment.
type Incrementer interface {
	Increment()
}