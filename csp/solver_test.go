package csp

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func setupAustraliaMapColoringCSP() *GenericCSP {
	states := []string{"WA", "NT", "SA", "Q", "NSW", "V", "T"}
	domains := make([]Domain, len(states))
	for i := range states {
		domains[i] = Domain{[]Value{1, 2, 3}}
	}
	constraints := []Constraint{
		NewAllDiffConstraint([]int{0, 1}),
		NewAllDiffConstraint([]int{0, 2}),
		NewAllDiffConstraint([]int{1, 2}),
		NewAllDiffConstraint([]int{1, 3}),
		NewAllDiffConstraint([]int{2, 3}),
		NewAllDiffConstraint([]int{2, 4}),
		NewAllDiffConstraint([]int{2, 5}),
		NewAllDiffConstraint([]int{3, 4}),
		NewAllDiffConstraint([]int{4, 5}),
	}

	return NewGenericCSP(domains, constraints)
}

func TestAustraliaMapColoringSolver(t *testing.T) {
	csp := setupAustraliaMapColoringCSP()
	solver := NewSimpleSolver(NextUnassignedVariableSelector, FirstDomainValueSelector)

	result := solver.Solve(csp)

	fmt.Println(result)

	assert.NotNil(t, result)
}

func BenchmarkAustraliaMapColoringSolver(b *testing.B) {
	csp := setupAustraliaMapColoringCSP()
	for i := 0; i < b.N; i++ {
		csp = NewGenericCSP(csp.domains, csp.constraints)
		solver := NewSimpleSolver(NextUnassignedVariableSelector, FirstDomainValueSelector)
		assert.NotNil(b, solver.Solve(csp))
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func setupNQueensCSP(n int) *GenericCSP {
	domains := make([]Domain, n)
	values := make([]Value, n)
	constraints := make([]Constraint, 0, n*(n-1))
	for i := 0; i < n; i++ {
		values[i] = Value(i)
	}
	for i := 0; i < n; i++ {
		domains[i] = Domain{slices.Clone(values)}
		for j := 0; j < n; j++ {
			if j != i {
				constraints = append(constraints, NewConstraint([]int{i, j}, func(indices []int, assignment Assignment) bool {
					return !assignment.Variables[indices[0]].Assigned ||
						!assignment.Variables[indices[1]].Assigned ||
						(assignment.Variables[indices[0]].Value != assignment.Variables[indices[1]].Value &&
							abs(int(assignment.Variables[indices[0]].Value)-int(assignment.Variables[indices[1]].Value)) != abs(i-j))
				}))
			}
		}
	}

	return NewGenericCSP(domains, constraints)
}

func TestNQueensSolver(t *testing.T) {
	csp := setupNQueensCSP(20)

	solver := NewSimpleSolver(NextUnassignedVariableSelector, FirstDomainValueSelector)

	result := solver.Solve(csp)

	fmt.Println(result)

	assert.NotNil(t, result)
}

func BenchmarkNQueensSolver(b *testing.B) {
	csp := setupNQueensCSP(20)
	for i := 0; i < b.N; i++ {
		csp = NewGenericCSP(csp.domains, csp.constraints)
		solver := NewSimpleSolver(NextUnassignedVariableSelector, FirstDomainValueSelector)
		assert.NotNil(b, solver.Solve(csp))
	}
}
