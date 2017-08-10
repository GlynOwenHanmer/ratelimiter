package ratelimiter_test

import (
	"testing"
	"github.com/GlynOwenHanmer/ratelimiter"
	"math"
	"time"
	"fmt"
)

func TestNewLimiter_NegativeRate(t *testing.T) {
	negativeRate := ratelimiter.Frequency(-10.0)
	limiter, err := ratelimiter.NewLimiter(negativeRate, 1)
	if err == nil {
		t.Error("Expected non-nil error but receieve nil.")
	}
	actualRate := limiter.Rate()
	if actualRate != 0 {
		t.Errorf("Expected limiter with rate of 0 for invalid NewLimiter paramater. Actual Rate is %f", actualRate)
	}
}

func TestNewLimiter_ZeroRate(t *testing.T) {
	zeroRate := ratelimiter.Frequency(0.0)
	limiter, err := ratelimiter.NewLimiter(zeroRate, 1)
	if err != nil {
		t.Errorf("Expected nil error but received: %s", err.Error())
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
	rate := ratelimiter.Frequency(1.0)
	limiter, err := ratelimiter.NewLimiter(rate, 1)
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

// A test to test the limiting accuracy of Limiters configured with various different rates.
func TestLimiter_Accuracy(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Limiter Allow test in short testing mode.")
	}
	// 2 ^ 15 = 32768, Highest rate will be 32768Hz
	for pow := 0; pow < 15; pow++ {
		rate := ratelimiter.Frequency(math.Pow(2,float64(pow)))
		limiter, err := ratelimiter.NewLimiter(rate, 1)
		if err != nil {
			t.Fatalf("Unexpected error creating limiter for testing: %s", err.Error())
		}
		allowCount := 0
		// 1 / 32768 = 30.5175781e-6 seconds, we need to oversample to have any chance of making an accurate test
		// So we sample the Limiter's Allow method appoximately 8 times for every time the token will be replenished
		ticker := time.NewTicker(time.Nanosecond * 30 / 8)
		go func() {
			for range ticker.C {
				if limiter.Allow() {
					allowCount++
				}
			}
		}()
		testLength := 3
		time.Sleep(time.Second * time.Duration(testLength))
		expectedAllowCount := int(rate) * testLength
		diff := math.Abs(float64(allowCount)-float64(expectedAllowCount))
		accuracy := 100.0 - diff / float64(expectedAllowCount) * 100
		t.Logf("Rate: %f, Allows: %d, Expected Allows: %d, Accuracy: %f%%", rate, allowCount, expectedAllowCount, accuracy)
	}
}

// Test that a burst can be achieved once the token bucket has been given time to fill.
func TestLimiter_Burst(t *testing.T) {
	rate := ratelimiter.Frequency(20)
	burst := uint64(10)
	limiter, err := ratelimiter.NewLimiter(rate, burst)
	if err != nil {
		t.Fatalf("Unexpected error creating limiter for testing: %s", err.Error())
	}
	// Allow limiter token bucket to fill
	time.Sleep(time.Millisecond * 500)
	for i := uint64(0); i < burst; i++ {
		if !limiter.Allow() {
			t.Errorf("Expected limiter to return true for Allow() but returned false.")
		}
	}
	if limiter.Allow() {
		t.Errorf("Expected limiter to return false after having Allow() called the maximum burst number of times")
	}
}

func ExampleNewLimiter() {
	rate := ratelimiter.Frequency(10)
	burst := uint64(3)
	limiter, err := ratelimiter.NewLimiter(rate, burst)
	if err != nil {
		fmt.Printf("Unable to create Limiter: %s", err.Error())
		return
	}
	printAllowed := func (allowed bool) {
		if allowed {
			// Your rate-limited event should be triggered here.
			fmt.Println("Allowed.")
			return
		}
		// Do not run your rate-limited event here.
		fmt.Println("Not allowed.")
	}
	// This sleep seems to be needed to hear to pass tests. Perhaps to give time for the Limiter Ticker to get started.
	time.Sleep(time.Millisecond * 10)
	printAllowed(limiter.Allow())
	time.Sleep(time.Millisecond * 100)
	printAllowed(limiter.Allow())
	printAllowed(limiter.Allow())
	printAllowed(limiter.Allow())
	time.Sleep(time.Millisecond * 400)
	printAllowed(limiter.Allow())
	printAllowed(limiter.Allow())
	printAllowed(limiter.Allow())
	printAllowed(limiter.Allow())
	printAllowed(limiter.Allow())
	// Output:
	// Allowed.
	// Allowed.
	// Not allowed.
	// Not allowed.
	// Allowed.
	// Allowed.
	// Allowed.
	// Not allowed.
	// Not allowed.
}