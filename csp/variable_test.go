package csp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVariable(t *testing.T) {
	v := Variable{}
	fmt.Println(v)

	v2 := v.Assign(3)
	assert.False(t, v.Assigned)
	assert.True(t, v2.Assigned)
	assert.Equal(t, v2.Value, 3)
}
