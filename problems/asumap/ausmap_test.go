package asumap

import (
	"fmt"
	"github.com/agrigoryan/gocsp/csp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAusMapSolving(t *testing.T) {
	problem := New()
	solver := csp.NewSimpleSolver(csp.NextUnassignedVariableSelector, csp.FirstDomainValueSelector, csp.NoInferenceFunc)

	result := solver.Solve(problem)

	fmt.Println(result)
	assert.NotNil(t, result)
}
