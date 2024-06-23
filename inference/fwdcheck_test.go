package inference

import (
	"fmt"
	"github.com/agrigoryan/gocsp/csp"
	"github.com/agrigoryan/gocsp/problems/asumap"
	"github.com/agrigoryan/gocsp/problems/nqueens"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAusMapWithFwdCheck(t *testing.T) {
	problem := asumap.New()
	solver := csp.NewBacktrackingSolver(csp.NextUnassignedVariableSelector, csp.FirstDomainValueSelector, FwdCheck)

	result := solver.Solve(problem)

	fmt.Println(result)
	assert.NotNil(t, result)
}

func BenchmarkAusMapWithFwdCheck(b *testing.B) {
	problem := asumap.New()
	benchWithFwdCheck(problem, b)
}

func TestNQueensWithFwdCheck(t *testing.T) {
	problem := nqueens.New(25)
	solver := csp.NewBacktrackingSolver(csp.NextUnassignedVariableSelector, csp.FirstDomainValueSelector, FwdCheck)

	result := solver.Solve(problem)

	fmt.Println(result)
	assert.NotNil(t, result)

}

func BenchmarkNQueensWithFwdCheck(b *testing.B) {
	problem := nqueens.New(25)
	benchWithFwdCheck(problem, b)
}

func benchWithFwdCheck(problem csp.CSP, b *testing.B) {
	for i := 0; i < b.N; i++ {
		problem = csp.NewGenericCSP(problem.Domains(), problem.Constraints())
		solver := csp.NewBacktrackingSolver(csp.NextUnassignedVariableSelector, csp.FirstDomainValueSelector, FwdCheck)
		assert.NotNil(b, solver.Solve(problem))
	}
}
