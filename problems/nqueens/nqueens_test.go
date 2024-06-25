package nqueens

import (
	"fmt"
	"github.com/agrigoryan/gocsp/csp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNQueensSolver(t *testing.T) {
	problem := New(25)

	solver := csp.NewBacktrackingSolver(csp.NextUnassignedVariableSelector, csp.FirstDomainValueSelector, nil)

	result := solver.Solve(problem)

	fmt.Println(result)
	assert.NotNil(t, result)
}
