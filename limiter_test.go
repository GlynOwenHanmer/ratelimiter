package PocketMediaLimiter_test

import (
	"testing"
	"github.com/GlynOwenHanmer/PocketMediaLimiter"
)

func TestNewLimiter_NegativeRate(t *testing.T) {
	negativeRate := -10.0
	limiter, err := PocketMediaLimiter.NewLimiter(negativeRate, 1)
	if err == nil {
		t.Error("Expected non-nil error but receieve nil.")
	}
	actualRate := limiter.Rate()
	if actualRate != 0 {
		t.Errorf("Expected limiter with rate of 0 for invalid NewLimiter paramater. Actual Rate is %f", actualRate)
	}
}

func TestNewLimiter_ZeroRate(t *testing.T) {
	zeroRate := 0.0
	limiter, err := PocketMediaLimiter.NewLimiter(zeroRate, 1)
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

func TestNewLimiter_PositiveRate(t *testing.T) {
	rate := 1.0
	limiter, err := PocketMediaLimiter.NewLimiter(rate, 1)
	if err != nil {
		t.Errorf("Expected nil error but received %s", err.Error())
	}
	actualRate := limiter.Rate()
	if actualRate != rate {
		t.Errorf("Expected rate of %f but got  %f", rate, actualRate)
	}
	if !limiter.Allow() {
		t.Errorf("Positive rate Limiter.Allow() should return true immeditaely after creating but returned true.")
	}
}

func TestLimiter_AllowAfterBurst(t *testing.T) {
	burst := uint64(5)
	limiter, err := PocketMediaLimiter.NewLimiter(1, burst)
	if err != nil {
		t.Fatalf("Unexpected error creating limiter for testing: %s", err.Error())
	}
	for i := uint64(0); i < burst; i++ {
		if !limiter.Allow() {
			t.Error("Limiter should Allow during burst just immediately after creating.")
		}
	}
	if limiter.Allow() {
		t.Error("Limiter should not Allow after having all tokens drained in burst")
	}
}
