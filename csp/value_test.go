package csp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewValueSet(t *testing.T) {
	s := NewValueSet([]Value{1, 2, 3, 4, 5})
	assert.Equal(t, 5, s.Size())
	assert.True(t, s.Contains(4))

	s = s.Add(44)
	assert.Equal(t, 6, s.Size())
	assert.True(t, s.Contains(44))
	assert.False(t, s.Contains(144))

	s = s.Remove(4)
	assert.Equal(t, 5, s.Size())
	assert.False(t, s.Contains(4))
}
