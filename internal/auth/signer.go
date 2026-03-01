package auth

import (
	"errors"
	"net/http"
	"strings"
)

const (
	headerAuthorization = "Authorization"
	headerWorkspace     = "X-DashScope-WorkSpace"
)

var (
	// ErrEmptyAPIKey indicates API key is missing.
	ErrEmptyAPIKey = errors.New("dashscope: API key is required")
)

// BuildHeaders builds websocket handshake headers for DashScope.
func BuildHeaders(apiKey, workspaceID string) (http.Header, error) {
	key := strings.TrimSpace(apiKey)
	if key == "" {
		return nil, ErrEmptyAPIKey
	}

	authorization := key
	if !strings.HasPrefix(strings.ToLower(authorization), "bearer ") {
		authorization = "bearer " + authorization
	}

	headers := make(http.Header)
	headers.Set(headerAuthorization, authorization)

	workspaceID = strings.TrimSpace(workspaceID)
	if workspaceID != "" {
		headers.Set(headerWorkspace, workspaceID)
	}

	return headers, nil
}
