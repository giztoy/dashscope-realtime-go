package dashscope

import (
	"net/http"
	"strings"
)

// Client is the DashScope SDK client.
type Client struct {
	Realtime *RealtimeService

	config *clientConfig
}

// NewClient creates a new DashScope client.
//
// The apiKey can be empty at construction time.
// Authentication errors are returned by Connect.
func NewClient(apiKey string, opts ...Option) *Client {
	apiKey = strings.TrimSpace(apiKey)

	cfg := &clientConfig{
		apiKey:            apiKey,
		baseURL:           DefaultRealtimeURL,
		httpBaseURL:       DefaultHTTPBaseURL,
		httpClient:        http.DefaultClient,
		connectTimeout:    defaultConnectTimeout,
		writeTimeout:      defaultWriteTimeout,
		readTimeout:       defaultReadTimeout,
		readLimitBytes:    defaultReadLimitBytes,
		reconnectAttempts: defaultReconnectAttempts,
		reconnectBase:     defaultReconnectBase,
		reconnectMax:      defaultReconnectMax,
	}

	for _, opt := range opts {
		if opt != nil {
			opt(cfg)
		}
	}

	if cfg.httpClient == nil {
		cfg.httpClient = http.DefaultClient
	}
	if cfg.connectTimeout < 0 {
		cfg.connectTimeout = defaultConnectTimeout
	}
	if cfg.writeTimeout <= 0 {
		cfg.writeTimeout = defaultWriteTimeout
	}
	if cfg.readTimeout < 0 {
		cfg.readTimeout = defaultReadTimeout
	}
	if cfg.readLimitBytes <= 0 {
		cfg.readLimitBytes = defaultReadLimitBytes
	}
	if cfg.reconnectAttempts < 0 {
		cfg.reconnectAttempts = 0
	}
	if cfg.reconnectBase <= 0 {
		cfg.reconnectBase = defaultReconnectBase
	}
	if cfg.reconnectMax < cfg.reconnectBase {
		cfg.reconnectMax = cfg.reconnectBase
	}

	c := &Client{config: cfg}
	c.Realtime = &RealtimeService{client: c}
	return c
}
