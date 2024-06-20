package inference

import (
	"fmt"
	"github.com/agrigoryan/gocsp/csp"
	"github.com/agrigoryan/gocsp/problems"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAusMapWithAC3(t *testing.T) {
	problem := problems.AusMap()
	solver := csp.NewSimpleSolver(csp.NextUnassignedVariableSelector, csp.FirstDomainValueSelector, AC3)

	result := solver.Solve(problem)

	fmt.Println(result)
	assert.NotNil(t, result)
}

func BenchmarkAusMapWithAC3(b *testing.B) {
	problem := problems.AusMap()
	benchWithAC3(problem, b)
}

func benchWithAC3(problem csp.CSP, b *testing.B) {
	for i := 0; i < b.N; i++ {
		problem = csp.NewGenericCSP(problem.Domains(), problem.Constraints())
		solver := csp.NewSimpleSolver(csp.NextUnassignedVariableSelector, csp.FirstDomainValueSelector, AC3)
		assert.NotNil(b, solver.Solve(problem))
	}
}
