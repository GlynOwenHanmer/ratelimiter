// internal tests to help with development of Limiter type whilst being able to access unexported fields and methods

package ratelimiter

import (
	"testing"
	"time"
)

func TestLimiter_ZeroRate_Allow(t *testing.T) {
	limiter := Limiter{rate:0.0,tokenBucket:tokenBucket{tokens:0,depth:100}}
	if limiter.Allow() {
		t.Errorf("Limiter a rate of 0 should return false Allow() with an empty bucket")
	}
	limiter.Increment()
	if !limiter.hasTokens(){
		t.Errorf("Expected hasTokens to return true but got false.")
	}
	if limiter.Allow() {
		t.Errorf("Limiter which has tokens and a rate of 0 should not return false for Allow()")
	}
}

func TestLimiter_Allow(t *testing.T) {
	limiter := Limiter{rate:1.0,tokenBucket:tokenBucket{tokens:0,depth:1}}
	if limiter.Allow() {
		t.Errorf("Limiter should not Allow() when it has an empty bucket.")
	}
	limiter.Increment()
	if !limiter.Allow() {
		t.Errorf("Expected Allow() to return true after being incremented to a single token in its tokenBucket.")
	}
	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			t.Errorf("Limiter should not Allow() after having Allow() called until empty.")
		}
	}
}

func TestLimiter_IsIncrementedOverTime(t *testing.T) {
	rate := Frequency(100.0)
	limiter, err := NewLimiter(rate, uint64(rate))
	if err != nil {
		t.Fatalf("Error when creating Limiter for testing. Error: %s", err.Error())
	}
	if limiter.tokens != 1 {
		t.Errorf("Limiter should start with 1 token.")
	}
	time.Sleep(time.Second)
	if limiter.tokens != limiter.tokenBucket.depth {
		t.Errorf("Expected Limiter tokenBucket to be full after 1 second. Expected %d tokens but got %d", limiter.tokenBucket.depth, limiter.tokens)
	}
}