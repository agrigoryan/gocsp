package csp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVariable(t *testing.T) {
	values := NewValueSet([]Value{1, 2, 3, 4, 5})
	domain := NewDomain(values)
	v := NewVariable("v1", domain)
	fmt.Println(v)

	assert.Equal(t, 5, v.Domain.Size())
	assert.False(t, v.Assigned)

	assert.Error(t, v.Assign(44))
	err := v.Assign(4)
	assert.Nil(t, err)
	assert.Equal(t, 1, v.Domain.Size())
	assert.True(t, v.Assigned)
	assert.Equal(t, Value(4), v.Value)
}
