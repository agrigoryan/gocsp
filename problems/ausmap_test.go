package problems

import (
	"fmt"
	"github.com/agrigoryan/gocsp/csp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAusMapSolving(t *testing.T) {
	problem := AusMap()
	solver := csp.NewSimpleSolver(csp.NextUnassignedVariableSelector, csp.FirstDomainValueSelector, csp.NoInferenceFunc)

	result := solver.Solve(problem)

	fmt.Println(result)

	assert.NotNil(t, result)
}

func BenchmarkAusMapSolving(b *testing.B) {
	problem := AusMap()
	for i := 0; i < b.N; i++ {
		problem = csp.NewGenericCSP(problem.Domains(), problem.Constraints())
		solver := csp.NewSimpleSolver(csp.NextUnassignedVariableSelector, csp.FirstDomainValueSelector, csp.NoInferenceFunc)
		assert.NotNil(b, solver.Solve(problem))
	}
}
