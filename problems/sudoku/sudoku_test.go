package sudoku

import (
	"fmt"
	"github.com/agrigoryan/gocsp/csp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSudokuSolver(t *testing.T) {
	input := `
1..4.8....5....7...........4.63.........7.59........2....8....16.......3.9..2....
`
	board := NewBoardFromString(input)
	fmt.Println(board)

	problem := New(board)
	solver := csp.NewBacktrackingSolver(csp.MRVVariableSelector, csp.FirstDomainValueSelector, InferenceFunc)
	result := solver.Solve(problem)

	assert.NotNil(t, result)

	resultBoard := NewBoard()
	for i := 0; i < len(result); i++ {
		resultBoard[i] = byte(result[i])
	}

	fmt.Println(resultBoard)
}
