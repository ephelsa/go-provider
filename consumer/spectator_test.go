package consumer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSpectator_Consume(t *testing.T) {
	// Given
	isCallbackCalled := false
	key := struct{ Key string }{Key: "TestKey"}
	spectator := NewSpectator(key, func(i int) {
		isCallbackCalled = true
		require.Equal(t, 10, i)
	})

	// When
	spectator.Consume(10)

	// Then
	assert.True(t, isCallbackCalled)
}

func TestSpectator_ConsumerKey(t *testing.T) {
	// Given
	key := struct{ Key string }{Key: "TestKey"}
	spectator := NewSpectator(key, func(i int) {
		// Not required
	})

	// When
	gotKey := spectator.ConsumerKey()

	// Then
	assert.Equal(t, key, gotKey)
}
