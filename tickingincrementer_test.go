package PocketMediaLimiter_test

import (
	"testing"
	"github.com/GlynOwenHanmer/PocketMediaLimiter"
	"time"
)

type mockIncrementer uint

func (mi *mockIncrementer) Increment() {
	*mi++
}

func TestNewTickingIncrementer_NegativeDuration(t *testing.T) {
	incrementer := mockIncrementer(0)
	ti, err := PocketMediaLimiter.NewTickingIncrementer(&incrementer, time.Second * -1)
	if err == nil {
		t.Errorf("Expected error but there wasn't one.")
	}
	ti.Stop()
}

func TestNewTickingIncrementer_NilIncrementer(t *testing.T) {
	ti, err := PocketMediaLimiter.NewTickingIncrementer(nil, time.Second * -1)
	if err == nil {
		t.Errorf("Expected error but there wasn't one.")
	}
	ti.Stop()
}

func TestNewTickingIncrementer(t *testing.T) {
	incrementer := mockIncrementer(0)
	duration := time.Millisecond * 50
	ti, err := PocketMediaLimiter.NewTickingIncrementer(&incrementer, duration)
	if err != nil {
		t.Fatalf("Error creating TickingIncrementor for testings. Error: %s", err.Error())
	}
	// Without this sleep, the timing isn't accurate enough to pass the tests.
	time.Sleep(time.Millisecond * 5)
	ticker := time.NewTicker(duration)
	expectedIncrements := uint64(0)
	go func() {
		for range ticker.C {
			expectedIncrements++
			actualIncrements := uint64(incrementer)
			if actualIncrements != expectedIncrements {
				t.Errorf("Expected %d increments but got %d", expectedIncrements, actualIncrements)
			}
		}
	}()
	time.Sleep(time.Millisecond * 500)
	ti.Stop()
	ticker.Stop()
	time.Sleep(time.Millisecond * 500)
	actualIncrements := uint64(incrementer)
	if actualIncrements != expectedIncrements {
		t.Errorf("Expected %d increments but got %d", expectedIncrements, actualIncrements)
	}
}