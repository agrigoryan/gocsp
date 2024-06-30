package sudoku

import (
	"fmt"
	"github.com/agrigoryan/gocsp/csp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSudokuSolver(t *testing.T) {
	input := "85...24..72......9..4.........1.7..23.5...9...4...........8..7..17..........36.4."
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
