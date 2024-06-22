package sudoku

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBoard(t *testing.T) {
	b := NewBoard()
	fmt.Println(b)
	assert.NotNil(t, b)
	assert.Equal(t, 81, len(b))
	assert.True(t, b.IsValid())
	assert.False(t, b.IsSolved())
}
