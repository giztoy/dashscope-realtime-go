package websocket

import (
	"context"
	"math"
	"time"
)

// BackoffPolicy controls reconnect delay.
type BackoffPolicy struct {
	BaseDelay time.Duration
	MaxDelay  time.Duration
	Factor    float64
}

func (p BackoffPolicy) normalized() BackoffPolicy {
	if p.BaseDelay <= 0 {
		p.BaseDelay = 200 * time.Millisecond
	}
	if p.MaxDelay <= 0 {
		p.MaxDelay = 2 * time.Second
	}
	if p.MaxDelay < p.BaseDelay {
		p.MaxDelay = p.BaseDelay
	}
	if p.Factor <= 1 {
		p.Factor = 2
	}
	return p
}

// Delay returns exponential backoff duration for a retry attempt (1-based).
func (p BackoffPolicy) Delay(attempt int) time.Duration {
	p = p.normalized()
	if attempt <= 1 {
		return p.BaseDelay
	}
	mul := math.Pow(p.Factor, float64(attempt-1))
	d := time.Duration(float64(p.BaseDelay) * mul)
	if d > p.MaxDelay {
		return p.MaxDelay
	}
	return d
}

// Retry retries fn with backoff until success, attempts exhausted, or context canceled.
func Retry(ctx context.Context, attempts int, policy BackoffPolicy, fn func(context.Context) error) error {
	if attempts <= 0 {
		attempts = 1
	}

	var lastErr error
	for i := 1; i <= attempts; i++ {
		if err := fn(ctx); err == nil {
			return nil
		} else {
			lastErr = err
		}

		if i == attempts {
			break
		}

		delay := policy.Delay(i)
		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		case <-timer.C:
		}
	}

	return lastErr
}
