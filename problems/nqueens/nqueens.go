package nqueens

import (
	"github.com/agrigoryan/gocsp/csp"
	"slices"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func New(n int) *csp.GenericCSP {
	domains := make([]csp.ValueSet, n)
	values := make([]csp.Value, n)
	constraints := make([]csp.Constraint, 0, n*(n-1))
	for i := 0; i < n; i++ {
		values[i] = csp.Value(i)
	}
	for i := 0; i < n; i++ {
		domains[i] = slices.Clone(values)
		for j := 0; j < n; j++ {
			if j != i {
				constraints = append(constraints, csp.NewConstraint([]int{i, j}, func(indices []int, assignment csp.Assignment) bool {
					v1, ok := assignment.AssignedValue(indices[0])
					if !ok {
						return true
					}
					v2, ok := assignment.AssignedValue(indices[1])
					if !ok {
						return true
					}
					return v1 != v2 && abs(int(v1)-int(v2)) != abs(i-j)
				}))
			}
		}
	}

	return csp.NewGenericCSP(domains, constraints)
}
