package asumap

import (
	"fmt"
	"github.com/agrigoryan/gocsp/csp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAusMapSolving(t *testing.T) {
	problem := New()
	solver := csp.NewBacktrackingSolver(csp.NextUnassignedVariableSelector, csp.FirstDomainValueSelector, nil)

	result := solver.Solve(problem)

	fmt.Println(result)
	assert.NotNil(t, result)
}
