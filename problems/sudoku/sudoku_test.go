package sudoku

import (
	"fmt"
	"github.com/agrigoryan/gocsp/csp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSudokuSolver(t *testing.T) {
	input := `
 *-----------*
 |.3.|4..|...|
 |9.2|8.6|3.1|
 |...|...|.2.|
 |---+---+---|
 |8..|.6.|7..|
 |.6.|2.5|.9.|
 |..3|.4.|..8|
 |---+---+---|
 |.7.|...|...|
 |4.8|9.2|5.6|
 |...|..8|.3.|
 *-----------*
`
	board := NewBoardFromString(input)
	fmt.Println(board)

	problem := New(board)
	solver := csp.NewSimpleSolver(csp.MRVVariableSelector, csp.FirstDomainValueSelector, csp.NoInferenceFunc)
	result := solver.Solve(problem)

	assert.NotNil(t, result)

	resultBoard := NewBoard()
	for i := 0; i < len(result); i++ {
		resultBoard[i] = byte(result[i])
	}

	fmt.Println(resultBoard)
}
