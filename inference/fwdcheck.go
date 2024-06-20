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
			if a.Variables[in].Assigned {
				continue
			}
			nDomain := a.Domain(in)
			for di := 0; di < nDomain.Size(); di++ {
				val := nDomain.Values()[di]
				a.Variable(in).Assign(val)
				if !c.IsSatisfied(assignment) {
					nDomain = nDomain.Remove(val)
					a.SetDomain(in, nDomain)
					di--
				}
			}
			a.Variable(in).Unassign()
			if nDomain.Size() == 0 {
				return assignment, false
			}
		}
	}
	return a, true
}
