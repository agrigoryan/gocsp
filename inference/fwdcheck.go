package inference

import "github.com/agrigoryan/gocsp/csp"

var FwdCheck csp.InferenceFunc = func(assignment csp.Assignment, constraints []csp.Constraint, varIdx int) (csp.Assignment, bool) {
	a := assignment.Copy()
	assignedVar := a.Variable(varIdx)
	for ic := 0; ic < len(assignedVar.Constraints); ic++ {
		c := assignedVar.Constraints[ic]
		neighborIndices := c.AppliesTo()
		for in := 0; in < len(neighborIndices); in++ {
			if in == varIdx {
				continue
			}
			neighborVar := a.Variable(in)
			if neighborVar.Assigned {
				continue
			}
			for di := 0; di < neighborVar.Domain.Size(); di++ {
				val := neighborVar.Domain.Value(di)
				neighborVar.Assign(val)
				if !c.IsSatisfied(assignment) {
					neighborVar.Domain.Remove(val)
					di--
				}
			}
			neighborVar.Unassign()
			if neighborVar.Domain.Size() == 0 {
				return assignment, false
			}
		}
	}
	return a, true
}
