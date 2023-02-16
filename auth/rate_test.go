package auth

import (
	"testing"
	"github.com/stretchr/testify/require"
)

// Test creating the rate manager
func TestCreateRateLimitManager(t *testing.T) {
	manager, err := CreateRateLimitManager(5, 1)
	require.Equal(t, err, nil)
	require.Equal(t, int(manager.maxPerSec), 5)
	require.Equal(t, manager.burst, 1)
	t.Log(manager)
}

// Test getting a visitor (new ip address)
func TestRateLimitNewVisitor(t *testing.T) {
	manager, _ := CreateRateLimitManager(5, 1)
	limit := manager.GetVisitor("test")
	require.Equal(t, limit != nil, true)
}

// Test fetching a visitor we have already created
func TestRateLimitOldVisitor(t *testing.T) {
	manager, _ := CreateRateLimitManager(5, 1)
	limit1 := manager.GetVisitor("test")
	limit2 := manager.GetVisitor("test")
	require.Equal(t, limit1, limit2)
}
