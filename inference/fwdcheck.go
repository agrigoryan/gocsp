package inference

import "github.com/agrigoryan/gocsp/csp"

var FwdCheck csp.InferenceFunc = func(assignment csp.Assignment, constraints []csp.Constraint, varIdx int) (csp.Assignment, bool) {
	a := assignment.Copy()
	assignedVar := a.Variables[varIdx]
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
			nDomain := a.Domains[in]
			for di := 0; di < nDomain.Size(); di++ {
				val := nDomain.Values()[di]
				a.Variables[in] = a.Variables[in].Assign(val)
				if !c.IsSatisfied(assignment) {
					// TODO: check how big is the performance difference if we use index access
					nDomain = nDomain.Remove(val)
					a.Domains[in] = nDomain
					di--
				}
			}
			a.Variables[in] = a.Variables[in].Unassign()
			if nDomain.Size() == 0 {
				return assignment, false
			}
		}
	}
	return a, true
}
