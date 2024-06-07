package csp

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveAustraliaMapColoring(t *testing.T) {
	states := []string{"WA", "NT", "SA", "Q", "NSW", "V", "T"}
	domains := make([]Domain, len(states))
	for i := range states {
		domains[i] = Domain{[]Value{1, 2, 3}}
	}
	constraints := []Constraint{
		AllDifferent{[]int{0, 1}},
		AllDifferent{[]int{0, 2}},
		AllDifferent{[]int{1, 2}},
		AllDifferent{[]int{1, 3}},
		AllDifferent{[]int{2, 3}},
		AllDifferent{[]int{2, 4}},
		AllDifferent{[]int{2, 5}},
		AllDifferent{[]int{3, 4}},
		AllDifferent{[]int{4, 5}},
	}

	csp := NewGenericCSP(domains, constraints)
	solver := NewSimpleSolver(NextUnassignedVariableSelector, FirstDomainValueSelector)

	result := solver.Solve(csp)

	fmt.Println(result)

	assert.NotNil(t, result)
}

func BenchmarkSolveAustraliaMapColoring(b *testing.B) {
	states := []string{"WA", "NT", "SA", "Q", "NSW", "V", "T"}
	domains := make([]Domain, len(states))
	for i := range states {
		domains[i] = Domain{[]Value{1, 2, 3}}
	}
	constraints := []Constraint{
		AllDifferent{[]int{0, 1}},
		AllDifferent{[]int{0, 2}},
		AllDifferent{[]int{1, 2}},
		AllDifferent{[]int{1, 3}},
		AllDifferent{[]int{2, 3}},
		AllDifferent{[]int{2, 4}},
		AllDifferent{[]int{2, 5}},
		AllDifferent{[]int{3, 4}},
		AllDifferent{[]int{4, 5}},
	}
	for i := 0; i < b.N; i++ {
		csp := NewGenericCSP(domains, constraints)
		solver := NewSimpleSolver(NextUnassignedVariableSelector, FirstDomainValueSelector)
		assert.NotNil(b, solver.Solve(csp))
	}
}
