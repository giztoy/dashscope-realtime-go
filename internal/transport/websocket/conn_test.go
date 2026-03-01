package websocket

import (
	"context"
	"errors"
	"io"
	"testing"
)

func TestIsRetryableClassification(t *testing.T) {
	if IsRetryable(nil) {
		t.Fatal("IsRetryable(nil) = true, want false")
	}
	if IsRetryable(context.Canceled) {
		t.Fatal("IsRetryable(context.Canceled) = true, want false")
	}
	if IsRetryable(ErrClosed) {
		t.Fatal("IsRetryable(ErrClosed) = true, want false")
	}
	if IsRetryable(ErrReconnectDisabled) {
		t.Fatal("IsRetryable(ErrReconnectDisabled) = true, want false")
	}
	if IsRetryable(errors.New("encode event: unsupported type")) {
		t.Fatal("IsRetryable(local encode error) = true, want false")
	}
	if !IsRetryable(io.EOF) {
		t.Fatal("IsRetryable(io.EOF) = false, want true")
	}
	if IsRetryable(&ConnectError{Err: errors.New("unauthorized"), StatusCode: 401}) {
		t.Fatal("IsRetryable(401 connect error) = true, want false")
	}
	if !IsRetryable(&ConnectError{Err: errors.New("server down"), StatusCode: 503}) {
		t.Fatal("IsRetryable(503 connect error) = false, want true")
	}
}

func TestNormalizeConfigReconnectAttemptsZeroIsPreserved(t *testing.T) {
	cfg := normalizeConfig(Config{ReconnectAttempts: 0})
	if cfg.ReconnectAttempts != 0 {
		t.Fatalf("ReconnectAttempts = %d, want 0", cfg.ReconnectAttempts)
	}
}

func TestReconnectDisabled(t *testing.T) {
	c := &Conn{cfg: Config{ReconnectAttempts: 0}}
	err := c.Reconnect(context.Background())
	if !errors.Is(err, ErrReconnectDisabled) {
		t.Fatalf("Reconnect() error = %v, want ErrReconnectDisabled", err)
	}
}
