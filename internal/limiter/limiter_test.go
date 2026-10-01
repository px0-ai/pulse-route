package limiter

import (
	"testing"
	"time"
)

func TestTokenBucketBasic(t *testing.T) {
	tb := NewTokenBucket(5, 1)

	for i := 0; i < 5; i++ {
		if !tb.Allow("127.0.0.1") {
			t.Fatalf("expected request %d to be allowed", i+1)
		}
	}

	if tb.Allow("127.0.0.1") {
		t.Fatal("expected request 6 to be rejected due to rate limit")
	}
}

func TestTokenBucketRefill(t *testing.T) {
	tb := NewTokenBucket(1, 10)
	if !tb.Allow("127.0.0.1") {
		t.Fatal("expected initial token to be allowed")
	}

	time.Sleep(150 * time.Millisecond)

	if !tb.Allow("127.0.0.1") {
		t.Fatal("expected refilled token to be allowed after delay")
	}
}
