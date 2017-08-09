// internal tests to help with development of Limiter type whilst being able to access unexported fields and methods

package PocketMediaLimiter

import (
	"testing"
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

//func TestLimiter_timedIncrementing(t *testing.T) {
//	rate := 50.0
//	limiter, _ := NewLimiter(rate, 100)
//	limiter.tokens = 0
//	 Without this sleep, the timing isn't accurate enough to pass the tests.
	//time.Sleep(time.Millisecond * 10)
	//ticker := createTicker(frequency(rate))
	//ticks := uint64(0)
	//go func() {
	//	for range ticker.C {
	//		ticks++
	//		actualTokens := limiter.tokens
	//		if actualTokens != ticks {
	//			t.Errorf("Expected %d tokens but has %d", ticks, actualTokens)
	//		}
	//	}
	//}()
	//time.Sleep(time.Second * 2)
	//ticker.Stop()
//}