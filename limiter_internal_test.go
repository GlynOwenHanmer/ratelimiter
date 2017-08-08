// internal tests to help with development of Limiter type whilst being able to access unexported fields and methods

package PocketMediaLimiter

import (
	"testing"
	"time"
)

func TestLimiter_increment(t *testing.T) {
	burst := uint64(25)
	limiter := Limiter{tokens:0, burst:burst}
	for i := uint64(0); i < burst; i++ {
		limiter.increment()
	}
	expectedTokens := uint64(burst)
	actualTokens := limiter.tokens
	if actualTokens != expectedTokens {
		t.Errorf("Expected %d tokens but got %d", expectedTokens, actualTokens)
	}
	for i := 0; uint64(i) < burst; i++ {
		limiter.increment()
	}
	expectedTokens = uint64(burst)
	actualTokens = limiter.tokens
	if actualTokens != expectedTokens {
		t.Errorf("Expected %d tokens but got %d", expectedTokens, actualTokens)
	}
}

func TestLimiter_decrement(t *testing.T) {
	burst := uint64(25)
	limiter := Limiter{tokens:burst, burst:burst}
	for i := uint64(0); i < burst; i++ {
		limiter.decrement()
	}
	expectedTokens := uint64(0)
	actualTokens := limiter.tokens
	if actualTokens != expectedTokens {
		t.Errorf("Expected %d tokens but got %d", expectedTokens, actualTokens)
	}
	for i := 0; uint64(i) < burst; i++ {
		limiter.decrement()
	}
	expectedTokens = uint64(0)
	actualTokens = limiter.tokens
	if actualTokens != expectedTokens {
		t.Errorf("Expected %d tokens but got %d", expectedTokens, actualTokens)
	}
}

func TestLimiter_incrementThenAllow(t *testing.T) {
	limiter := Limiter{rate:1.0,tokens:0, burst:1}
	limiter.increment()
	if !limiter.hasTokens() {
		t.Errorf("Expected limiter.hasTokens to be true after incremenet but got false")
	}
	if !limiter.Allow() {
		t.Errorf("Limiter should Allow() after having been incremented.")
	}
	if limiter.Allow() {
		t.Errorf("Limiter should not Allow() after having been incremented then Allowed.")
	}
}

func TestLimiter_incrementToBurstLimit(t *testing.T) {
	limiter := Limiter{rate:1.0, tokens:0, burst:1}
	limiter.increment()
	limiter.increment()
	expectedTokens := uint64(1)
	actualTokens := limiter.tokens
	if actualTokens != expectedTokens {
		t.Errorf("Expected %d tokens but got %d", expectedTokens, actualTokens)
	}
	if !limiter.hasTokens() {
		t.Errorf("Expected limiter.tokens to be true after incremenet but got false")
	}
	if !limiter.Allow() {
		t.Errorf("Limiter should Allow() after having been incremented twice.")
	}
	expectedTokens = uint64(0)
	actualTokens = limiter.tokens
	if actualTokens != expectedTokens {
		t.Errorf("Expected %d tokens but got %d", expectedTokens, actualTokens)
	}
	if limiter.Allow() {
		t.Errorf("Limiter should not Allow() after having been incremented twice then Allowed.")
	}
}

func TestLimiter_decrementFromZeroTokens(t *testing.T) {
	limiter := Limiter{rate:1.0, tokens:0, burst:1}
	limiter.decrement()
	expectedTokens := uint64(0)
	actualTokens := limiter.tokens
	if actualTokens != expectedTokens {
		t.Errorf("Expected %d tokens but got %d", expectedTokens, actualTokens)
	}
	if limiter.hasTokens() {
		t.Error("Expected limiter.tokens to be false after decrement but got true")
	}
	if limiter.Allow() {
		t.Error("Limiter.Allow() should not return true after having been decremented to 0 tokens.")
	}
	limiter.decrement()
	expectedTokens = uint64(0)
	actualTokens = limiter.tokens
	if actualTokens != expectedTokens {
		t.Errorf("Expected %d tokens but got %d", expectedTokens, actualTokens)
	}
	if limiter.hasTokens() {
		t.Error("Expected limiter.tokens to be false after decrement but got true")
	}
	if limiter.Allow() {
		t.Error("Limiter.Allow() should not return true after having been decremented to 0 tokens.")
	}
}

func TestLimiter_timedIncrementing(t *testing.T) {
	rate := 50.0
	limiter, _ := NewLimiter(rate, 100)
	limiter.tokens = 0
	// Without this sleep, the timing isn't accurate enough to pass the tests.
	time.Sleep(time.Millisecond * 10)
	ticker := createTicker(frequency(rate))
	ticks := uint64(0)
	go func() {
		for range ticker.C {
			ticks++
			actualTokens := limiter.tokens
			if actualTokens != ticks {
				t.Errorf("Expected %d tokens but has %d", ticks, actualTokens)
			}
		}
	}()
	timer := time.NewTimer(time.Second * 2)
	<- timer.C
}