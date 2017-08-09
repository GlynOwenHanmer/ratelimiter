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

//A test to test the limiting accuracy of Limiters configured with various different rates.
//func TestLimiter_Accuracy(t *testing.T) {
//	if testing.Short() {
//		t.Skip("Skipping Limiter Allow test in short testing mode.")
//	}
//	 2 ^ 15 = 32768, Highest rate will be 32768Hz
	//for pow := 0; pow < 15; pow++ {
	//	rate := math.Pow(2,float64(pow))
	//	limiter, err := PocketMediaLimiter.NewLimiter(rate, 1)
	//	if err != nil {
	//		t.Fatalf("Unexpected error creating limiter for testing: %s", err.Error())
	//	}
	//	allowCount := 0
	//	 1 / 32768 = 30.5175781e-6 seconds, we need to oversample to have any chance of making an accurate test
	//	 So we sample the Limiter's Allow method appoximately 8 times for every time the token will be replenished
		//ticker := time.NewTicker(time.Nanosecond * 30 / 8)
		//go func() {
		//	for range ticker.C {
		//		if limiter.Allow() {
		//			allowCount++
		//		}
		//	}
		//}()
		//testLength := 3
		//time.Sleep(time.Second * time.Duration(testLength))
		//expectedAllowCount := int(rate) * testLength
		//diff := math.Abs(float64(allowCount)-float64(expectedAllowCount))
		//accuracy := 100.0 - diff / float64(expectedAllowCount) * 100
		//t.Logf("Rate: %f, Allows: %d, Expected Allows: %d, Accuracy: %f", rate, allowCount, expectedAllowCount, accuracy)
	//}
//}
