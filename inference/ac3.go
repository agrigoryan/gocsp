package inference

import "github.com/agrigoryan/gocsp/csp"

type arc struct {
	s, t int
	c    csp.Constraint
}

var AC3 csp.InferenceFunc = func(assignment csp.Assignment, constraints []csp.Constraint, varIdx int) (csp.Assignment, bool) {
	as := assignment.Copy()
	//constraints = as.Variables[varIdx].Constraints

	// initially populate the queue with the arcs of the newly assigned variable
	queue := make([]arc, 0, 2*len(constraints))
	for _, c := range constraints {
		cIndices := c.AppliesTo()
		if c.IsBinaryConstraint() && !as.Variables[cIndices[0]].Assigned {
			queue = append(queue, arc{s: cIndices[0], t: cIndices[1], c: c}, arc{s: cIndices[1], t: cIndices[0], c: c})
		}
	}

	revise := func(a arc) bool {
		idx1, idx2 := a.s, a.t
		v1, v2 := as.Variables[idx1], as.Variables[idx2]
		d1, d2 := as.Domains[idx1], as.Domains[idx2]

		revised := false

		var di, dj csp.Value
		for i := 0; i < d1.Size(); i++ {
			di = d1.Values()[i]
			as.Variables[idx1] = v1.Assign(di)
			anyValueSatisfiesArc := false
			if v2.Assigned {
				anyValueSatisfiesArc = a.c.IsSatisfied(as)
			} else {
				for j := 0; j < d2.Size(); j++ {
					dj = d2.Values()[j]
					as.Variables[idx2] = v2.Assign(dj)
					arcSatisfied := a.c.IsSatisfied(as)
					as.Variables[idx2] = v2
					if arcSatisfied {
						anyValueSatisfiesArc = true
						break
					}
				}
			}
			if !anyValueSatisfiesArc {
				d1 = d1.Remove(di)
				as.Domains[idx1] = d1
				i--
				revised = true
			}
		}

		as.Variables[idx1] = v1
		as.Variables[idx2] = v2

		return revised
	}

	for len(queue) > 0 {
		a := queue[0]
		queue = queue[1:]
		varIdx := a.s

		if revise(a) {
			if as.Domains[varIdx].Size() == 0 {
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
				if c.IsBinaryConstraint() && neighborIdx != a.t {
					queue = append(queue, arc{s: neighborIdx, t: varIdx, c: c})
				}
			}
		}
	}

	return as, true
}
