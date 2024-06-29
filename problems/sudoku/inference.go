package sudoku

import "github.com/agrigoryan/gocsp/csp"

var InferenceFunc csp.InferenceFunc = func(a *csp.Assignment, constraints []csp.Constraint, varIdx int) bool {
	vi, _ := a.AssignedValue(varIdx)
	for _, c := range a.Constraints(varIdx) {
		for _, i := range c.AppliesTo() {
			if i == varIdx || a.Assigned(i) {
				continue
			}
			a.FilterDomain(i, func(valueIdx int) bool {
				return a.DomainValue(i, valueIdx) != vi
			})
			if a.DomainSize(i) == 0 {
				return false
			}
		}
	}
	return true
}
