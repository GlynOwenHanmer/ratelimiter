package PocketMediaLimiter

import "testing"

func TestLimiter_increment(t *testing.T) {
	limiter := Limiter{tokens:false}
	limiter.increment()
	if !limiter.tokens {
		t.Errorf("Expected limiter.tokens to be true after incremenet but got false")
	}
}

func TestLimiter_decrement(t *testing.T) {
	limiter := Limiter{tokens:true}
	limiter.decrement()
	if limiter.tokens {
		t.Errorf("Expected limiter.tokens to be false after decremenet but got true")
	}
}

func TestLimiter_incrementThenAllow(t *testing.T) {
	limiter := Limiter{rate:1.0,tokens:false}
	limiter.increment()
	if !limiter.tokens {
		t.Errorf("Expected limiter.tokens to be true after incremenet but got false")
	}
	if !limiter.Allow() {
		t.Errorf("Limiter should Allow() after having been incremented.")
	}
	if limiter.Allow() {
		t.Errorf("Limiter should not Allow() after having been incremented then Allowed.")
	}
}

func TestLimiter_incrementToLimit(t *testing.T) {
	limiter := Limiter{rate:1.0, tokens:false}
	limiter.increment()
	limiter.increment()
	if !limiter.tokens {
		t.Errorf("Expected limiter.tokens to be true after incremenet but got false")
	}
	if !limiter.Allow() {
		t.Errorf("Limiter should Allow() after having been incremented twice.")
	}
	if limiter.Allow() {
		t.Errorf("Limiter should not Allow() after having been incremented twice then Allowed.")
	}
}