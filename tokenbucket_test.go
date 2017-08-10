package ratelimiter

import (
	"testing"
)

func TestTokenBucket_Increment(t *testing.T) {
	depth := uint64(25)
	bucket := tokenBucket{tokens:0, depth:depth}
	for i := uint64(0); i < depth; i++ {
		bucket.Increment()
		expectedTokens := i + 1
		actualTokens := bucket.tokens
		if actualTokens != expectedTokens {
			t.Errorf("Expected %d tokens but got %d whilst filling up bucket.", expectedTokens, actualTokens)
		}
	}
	expectedTokens := uint64(depth)
	actualTokens := bucket.tokens
	if actualTokens != expectedTokens {
		t.Errorf("Expected %d tokens but got %d after filling bucket.", expectedTokens, actualTokens)
	}
	for i := 0; uint64(i) < depth; i++ {
		bucket.Increment()
	}
	expectedTokens = uint64(depth)
	actualTokens = bucket.tokens
	if actualTokens != expectedTokens {
		t.Errorf("Expected %d tokens but got %d after attempting to overfill bucket", expectedTokens, actualTokens)
	}
}

func TestTokenBucket_decrement(t *testing.T) {
	depth := uint64(25)
	bucket := tokenBucket{tokens:depth, depth:depth}
	for i := uint64(0); i < depth; i++ {
		bucket.decrement()
		expectedTokens := depth - i - 1
		actualTokens := bucket.tokens
		if actualTokens != expectedTokens {
			t.Errorf("Expected %d tokens but got %d whilst emptying bucket.", expectedTokens, actualTokens)
		}
	}
	expectedTokens := uint64(0)
	actualTokens := bucket.tokens
	if actualTokens != expectedTokens {
		t.Errorf("Expected %d tokens but got %d having emptied bucket", expectedTokens, actualTokens)
	}
	for i := 0; uint64(i) < depth; i++ {
		bucket.decrement()
	}
	expectedTokens = uint64(0)
	actualTokens = bucket.tokens
	if actualTokens != expectedTokens {
		t.Errorf("Expected %d tokens but got %d after attempting to further empty bucket", expectedTokens, actualTokens)
	}
}

func TestTokenBucket_HasTokens(t *testing.T) {
	testSets := []struct{
		tokens uint64
		bool
	}{
		{0,false},
		{1, true},
	}
	for _, testSet := range testSets {
		bucket := tokenBucket{tokens:testSet.tokens}
		hasTokens := bucket.hasTokens()
		if hasTokens != testSet.bool {
			t.Errorf("Unexpected hasTokens() result.\n\tExpected: %t\n\tActual  : %t", hasTokens, testSet.bool)
		}
	}
}
