package PocketMediaLimiter

// tokenBucket holds a number representing a count of tokens within a bucket with a given depth.
// A tokenBucket can have tokens added to it and taken away from it but will never overflow the depth of the bucket or have a negative value for the number of tokens it holds.
type tokenBucket struct {
	tokens uint64
	depth uint64
}

// increment increments the number of tokens available in the bucket if there is space in the bucket to do so.
func (b *tokenBucket) Increment() {
	if b.tokens >= b.depth {
		return
	}
	b.tokens++
}

// decrement decrements the number of tokens available in the bucket if there are any left.
func (b *tokenBucket) decrement() {
	if b.tokens == 0 {
		return
	}
	b.tokens--
}

// hasTokens returns true if the bucket has available tokens.
func (b *tokenBucket) hasTokens() bool {
	return b.tokens > 0
}