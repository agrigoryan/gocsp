package inference

import "github.com/agrigoryan/gocsp/csp"

var FwdCheck csp.InferenceFunc = func(assignment csp.Assignment, constraints []csp.Constraint, varIdx int) (csp.Assignment, bool) {
	assignedVar := assignment.Variable(varIdx)
	for ic := 0; ic < len(assignedVar.Constraints); ic++ {
		c := assignedVar.Constraints[ic]
		neighborIndices := c.AppliesTo()
		for in := 0; in < len(neighborIndices); in++ {
			if in == varIdx {
				continue
			}
			neighborVar := assignment.Variable(in)
			if neighborVar.Assigned {
				continue
			}
			neighborVar.Domain.Filter(func(idx int) bool {
				neighborVar.Assign(idx)
				return c.IsSatisfied(assignment)
			})
			neighborVar.Unassign()
			if neighborVar.Domain.Size() == 0 {
				return assignment, false
			}
		}
	}
	return assignment, true
}
