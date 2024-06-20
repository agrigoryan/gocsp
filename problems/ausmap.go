package problems

import "github.com/agrigoryan/gocsp/csp"

func AusMap() *csp.GenericCSP {
	states := []string{"WA", "NT", "SA", "Q", "NSW", "V", "T"}
	domains := make([]csp.Domain, len(states))
	for i := range states {
		domains[i] = csp.NewDomain([]csp.Value{1, 2, 3})
	}
	constraints := []csp.Constraint{
		csp.NewAllDiffConstraint([]int{0, 1}),
		csp.NewAllDiffConstraint([]int{0, 2}),
		csp.NewAllDiffConstraint([]int{1, 2}),
		csp.NewAllDiffConstraint([]int{1, 3}),
		csp.NewAllDiffConstraint([]int{2, 3}),
		csp.NewAllDiffConstraint([]int{2, 4}),
		csp.NewAllDiffConstraint([]int{2, 5}),
		csp.NewAllDiffConstraint([]int{3, 4}),
		csp.NewAllDiffConstraint([]int{4, 5}),
	}

	return csp.NewGenericCSP(domains, constraints)
}