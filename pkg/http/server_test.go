package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListenAddress(t *testing.T) {
	tests := []struct {
		name        string
		bindAddress string
		port        int
		expected    string
	}{
		{
			name:     "all interfaces",
			port:     8082,
			expected: ":8082",
		},
		{
			name:        "IPv4 loopback",
			bindAddress: "127.0.0.1",
			port:        8082,
			expected:    "127.0.0.1:8082",
		},
		{
			name:        "IPv6 loopback",
			bindAddress: "::1",
			port:        8082,
			expected:    "[::1]:8082",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := listenAddress(tt.bindAddress, tt.port)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
