package dashscope

import (
	"testing"
	"time"
)

func TestWithReconnectZeroDisablesReconnect(t *testing.T) {
	client := NewClient("test-key", WithReconnect(0, 100*time.Millisecond, 1*time.Second))
	if client.config.reconnectAttempts != 0 {
		t.Fatalf("reconnectAttempts = %d, want 0", client.config.reconnectAttempts)
	}
}
