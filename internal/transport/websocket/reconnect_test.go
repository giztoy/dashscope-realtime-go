package websocket

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestBackoffPolicyDelay(t *testing.T) {
	policy := BackoffPolicy{
		BaseDelay: 100 * time.Millisecond,
		MaxDelay:  500 * time.Millisecond,
		Factor:    2,
	}

	tests := []struct {
		attempt int
		want    time.Duration
	}{
		{attempt: 1, want: 100 * time.Millisecond},
		{attempt: 2, want: 200 * time.Millisecond},
		{attempt: 3, want: 400 * time.Millisecond},
		{attempt: 4, want: 500 * time.Millisecond},
	}

	for _, tt := range tests {
		if got := policy.Delay(tt.attempt); got != tt.want {
			t.Fatalf("Delay(%d) = %v, want %v", tt.attempt, got, tt.want)
		}
	}
}

func TestRetryEventuallySucceeds(t *testing.T) {
	var count int
	err := Retry(context.Background(), 5, BackoffPolicy{BaseDelay: 1 * time.Millisecond, MaxDelay: 1 * time.Millisecond}, func(context.Context) error {
		count++
		if count < 3 {
			return errors.New("temporary")
		}
		return nil
	})

	if err != nil {
		t.Fatalf("Retry() error = %v", err)
	}
	if count != 3 {
		t.Fatalf("Retry() attempts = %d, want 3", count)
	}
}
