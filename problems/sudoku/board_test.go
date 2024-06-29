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

func TestNewBoardFromString(t *testing.T) {
	tests := []struct {
		board  string
		valid  bool
		solved bool
	}{
		{
			board:  "7.8...3.....6.1...5.........4.....263...8.......1...9..9.2....4....7.5...........",
			valid:  true,
			solved: false,
		},
		{
			board:  "468931527751624839392578461134756298289413675675289314846192753513867942927345186",
			valid:  true,
			solved: true,
		},
		{
			board:  "7.8...3.....6.1...5.........4..",
			valid:  false,
			solved: false,
		},
		{
			board: `
*---*---*---*
|4..|.3.|...|
|...|6..|8..|
|...|...|..1|
*---*---*---*
|...|.5.|.9.|
|.8.|...|6..|
|.7.|2..|...|
*---*---*---*
|...|1.2|7..|
|5.3|...|.4.|
|9..|...|...|
*---*---*---*`,
			valid:  true,
			solved: false,
		},
	}

	for _, test := range tests {
		t.Run(test.board, func(t *testing.T) {
			b := NewBoardFromString(test.board)
			fmt.Println(b)
			assert.NotNil(t, b)
			assert.Equal(t, test.valid, b.IsValid())
			assert.Equal(t, test.solved, b.IsSolved())
		})
	}
}
