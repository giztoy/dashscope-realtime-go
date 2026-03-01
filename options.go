package dashscope

import (
	"net/http"
	"time"
)

const (
	// DefaultRealtimeURL is the default DashScope Realtime websocket endpoint.
	DefaultRealtimeURL = "wss://dashscope.aliyuncs.com/api-ws/v1/realtime"

	// DefaultHTTPBaseURL is the default DashScope HTTP endpoint.
	DefaultHTTPBaseURL = "https://dashscope.aliyuncs.com"

	defaultConnectTimeout = 15 * time.Second
	defaultWriteTimeout   = 10 * time.Second
	defaultReadTimeout    = 0
	defaultReadLimitBytes = int64(8 << 20) // 8MiB

	defaultReconnectAttempts = 3
	defaultReconnectBase     = 200 * time.Millisecond
	defaultReconnectMax      = 2 * time.Second
)

// Option configures a Client.
type Option func(*clientConfig)

type clientConfig struct {
	apiKey      string
	workspaceID string

	baseURL     string
	httpBaseURL string
	httpClient  *http.Client

	connectTimeout time.Duration
	writeTimeout   time.Duration
	readTimeout    time.Duration
	readLimitBytes int64

	reconnectAttempts int
	reconnectBase     time.Duration
	reconnectMax      time.Duration
}

// WithWorkspace sets workspace ID for request isolation.
func WithWorkspace(workspaceID string) Option {
	return func(c *clientConfig) {
		c.workspaceID = workspaceID
	}
}

// WithBaseURL sets websocket base URL.
func WithBaseURL(url string) Option {
	return func(c *clientConfig) {
		c.baseURL = url
	}
}

// WithHTTPBaseURL sets HTTP base URL.
func WithHTTPBaseURL(url string) Option {
	return func(c *clientConfig) {
		c.httpBaseURL = url
	}
}

// WithHTTPClient sets custom HTTP client.
func WithHTTPClient(client *http.Client) Option {
	return func(c *clientConfig) {
		if client != nil {
			c.httpClient = client
		}
	}
}

// WithConnectTimeout sets websocket connect timeout.
func WithConnectTimeout(timeout time.Duration) Option {
	return func(c *clientConfig) {
		c.connectTimeout = timeout
	}
}

// WithReadTimeout sets per-read timeout.
// Zero means no additional read deadline.
func WithReadTimeout(timeout time.Duration) Option {
	return func(c *clientConfig) {
		c.readTimeout = timeout
	}
}

// WithReadLimitBytes sets max websocket message size in bytes.
func WithReadLimitBytes(limit int64) Option {
	return func(c *clientConfig) {
		c.readLimitBytes = limit
	}
}

// WithWriteTimeout sets per-write timeout.
func WithWriteTimeout(timeout time.Duration) Option {
	return func(c *clientConfig) {
		c.writeTimeout = timeout
	}
}

// WithReconnect configures reconnect attempts and backoff window.
// attempts=0 disables reconnect.
func WithReconnect(attempts int, baseDelay, maxDelay time.Duration) Option {
	return func(c *clientConfig) {
		c.reconnectAttempts = attempts
		c.reconnectBase = baseDelay
		c.reconnectMax = maxDelay
	}
}
