package PocketMediaLimiter_test

import (
	"testing"
	"github.com/GlynOwenHanmer/PocketMediaLimiter"
)

func TestNewLimiter_NegativeRate(t *testing.T) {
	negativeRate := -10.0
	limiter, err := PocketMediaLimiter.NewLimiter(negativeRate)
	if err == nil {
		t.Error("Expected non-nil error but receieve nil.")
	}
	actualRate := limiter.Rate()
	if actualRate != 0 {
		t.Errorf("Expected limiter with rate of 0 for invalid NewLimiter paramater. Actual Rate is %f", actualRate)
	}
}

//todo Add code to test for when Zero Rate Limiter bucket has been replenished/incremented
func TestNewLimiter_ZeroRate(t *testing.T) {
	zeroRate := 0.0
	limiter, err := PocketMediaLimiter.NewLimiter(zeroRate)
	if err != nil {
		t.Errorf("Expected nil error but received %s", err.Error())
	}
	actualRate := limiter.Rate()
	if actualRate != zeroRate {
		t.Errorf("Expected rate of %f but got  %f", zeroRate, actualRate)
	}
	if limiter.Allow() {
		t.Errorf("Zero rate Limiter.Allow() should return false but returned true.")
	}
}