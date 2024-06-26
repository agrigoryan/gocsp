package inference

import "github.com/agrigoryan/gocsp/csp"

var FwdCheck csp.InferenceFunc = func(assignment csp.Assignment, constraints []csp.Constraint, varIdx int) (csp.Assignment, bool) {
	for _, c := range assignment.Constraints(varIdx) {
		neighborIndices := c.AppliesTo()
		for in := 0; in < len(neighborIndices); in++ {
			if in == varIdx {
				continue
			}
			if assignment.Assigned(in) {
				continue
			}
			assignment.FilterDomain(in, func(idx int) bool {
				assignment.Assign(in, idx)
				return c.IsSatisfied(assignment)
			})
			assignment.Unassign(in)
			if assignment.DomainSize(in) == 0 {
				return assignment, false
			}
		}
	}
	return assignment, true
}
