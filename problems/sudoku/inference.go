package sudoku

import "github.com/agrigoryan/gocsp/csp"

var InferenceFunc csp.InferenceFunc = func(a *csp.Assignment, constraints []csp.Constraint, varIdx int) bool {
	toCheck := []int{varIdx}
	counter := 0

	for len(toCheck) > 0 {
		varIdx = toCheck[0]
		toCheck = toCheck[1:]
		counter++
		vi, _ := a.AssignedValue(varIdx)
		for _, c := range a.Constraints(varIdx) {
			for _, i := range c.AppliesTo() {
				if i == varIdx || a.Assigned(i) {
					continue
				}
				a.FilterDomain(i, func(valueIdx int) bool {
					return a.DomainValue(i, valueIdx) != vi
				})
				domainSize := a.DomainSize(i)
				if domainSize == 0 {
					return false
				}
				if domainSize == 1 {
					a.RangeDomain(i, func(valueIdx int) bool {
						a.Assign(i, valueIdx)
						toCheck = append(toCheck, i)
						return true
					})
				}
			}
		}
	}

	// TODO: check if we absolutely need to check the constraints again
	if counter < 2 {
		return true
	}
	for _, c := range constraints {
		if !c.IsSatisfied(a) {
			return false
		}
	}

	return true
}
