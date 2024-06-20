package inference

import "github.com/agrigoryan/gocsp/csp"

type arc struct {
	s, t int
	c    csp.Constraint
}

func ac3Revise(assignment csp.Assignment, a arc) bool {
	idx1, idx2 := a.s, a.t
	v1, v2 := assignment.Variable(idx1), assignment.Variable(idx2)
	d1, d2 := v1.Domain, v2.Domain

	revised := false

	var di, dj csp.Value
	for i := 0; i < d1.Size(); i++ {
		di = d1.Values()[i]
		v1.Assign(di)
		anyValueSatisfiesArc := false
		if v2.Assigned {
			anyValueSatisfiesArc = a.c.IsSatisfied(assignment)
		} else {
			for j := 0; j < d2.Size(); j++ {
				dj = d2.Values()[j]
				v2.Assign(dj)
				arcSatisfied := a.c.IsSatisfied(assignment)
				v2.Unassign()
				if arcSatisfied {
					anyValueSatisfiesArc = true
					break
				}
			}
		}
		if !anyValueSatisfiesArc {
			v1.RemoveFromDomain(di)
			revised = true
		}
	}

	v1.Unassign()

	return revised
}

var AC3 csp.InferenceFunc = func(assignment csp.Assignment, constraints []csp.Constraint, varIdx int) (csp.Assignment, bool) {
	as := assignment.Copy()
	//constraints = as.Variables[varIdx].Constraints

	// initially populate the queue with the arcs of the newly assigned variable
	queue := make([]arc, 0, 2*len(constraints))
	for _, c := range constraints {
		cIndices := c.AppliesTo()
		if c.IsBinaryConstraint() && !as.Variable(cIndices[0]).Assigned {
			queue = append(queue, arc{s: cIndices[0], t: cIndices[1], c: c}, arc{s: cIndices[1], t: cIndices[0], c: c})
		}
	}

	for len(queue) > 0 {
		a := queue[0]
		queue = queue[1:]
		varIdx := a.s
		variable := as.Variable(varIdx)

		if ac3Revise(as, a) {
			if variable.Domain.Size() == 0 {
				return assignment, false
			}

			var c csp.Constraint
			for i := 0; i < len(as.Variables[varIdx].Constraints); i++ {
				c = as.Variables[varIdx].Constraints[i]
				cIndices := c.AppliesTo()
				neighborIdx := cIndices[0]
				if neighborIdx == varIdx {
					neighborIdx = cIndices[1]
				}
				if neighborIdx != a.t && c.IsBinaryConstraint() && !as.Variable(neighborIdx).Assigned {
					queue = append(queue, arc{s: neighborIdx, t: varIdx, c: c})
				}
			}
		}
	}

	return as, true
}
