package inference

import (
	"fmt"
	"github.com/agrigoryan/gocsp/csp"
	"github.com/agrigoryan/gocsp/problems"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAusMapWithFwdCheck(t *testing.T) {
	problem := problems.AusMap()
	solver := csp.NewSimpleSolver(csp.NextUnassignedVariableSelector, csp.FirstDomainValueSelector, FwdCheck)

	result := solver.Solve(problem)

	fmt.Println(result)
	assert.NotNil(t, result)
}

func BenchmarkAusMapWithFwdCheck(b *testing.B) {
	problem := problems.AusMap()
	benchWithFwdCheck(problem, b)
}

func TestNQueensWithFwdCheck(t *testing.T) {
	problem := problems.NQueens(25)
	solver := csp.NewSimpleSolver(csp.NextUnassignedVariableSelector, csp.FirstDomainValueSelector, FwdCheck)

	result := solver.Solve(problem)

	fmt.Println(result)
	assert.NotNil(t, result)

}

func BenchmarkNQueensWithFwdCheck(b *testing.B) {
	problem := problems.NQueens(25)
	benchWithFwdCheck(problem, b)
}

func benchWithFwdCheck(problem csp.CSP, b *testing.B) {
	for i := 0; i < b.N; i++ {
		problem = csp.NewGenericCSP(problem.Domains(), problem.Constraints())
		solver := csp.NewSimpleSolver(csp.NextUnassignedVariableSelector, csp.FirstDomainValueSelector, FwdCheck)
		assert.NotNil(b, solver.Solve(problem))
	}
}
