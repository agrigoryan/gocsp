package inference

import "github.com/agrigoryan/gocsp/csp"

type arc struct {
	x, y int
	c    csp.Constraint
	next *arc
}

type arcQueue struct {
	head *arc
	tail *arc
}

func (q *arcQueue) push(a *arc) {
	if q.head == nil {
		q.head = a
	}
	if q.tail != nil {
		q.tail.next = a
	}
	q.tail = a
}

func (q *arcQueue) shift() *arc {
	if q.head == nil {
		return nil
	}
	a := q.head
	q.head = a.next
	if q.tail == a {
		q.tail = nil
	}
	return a
}

func (q *arcQueue) empty() bool {
	return q.head == nil
}

func ac3Revise(assignment csp.Assignment, a *arc) bool {
	vx, vy := assignment.Variable(a.x), assignment.Variable(a.y)
	if vx.Assigned {
		panic("BOOO")
	}

	revised := false

	var di, dj csp.Value
	for i := 0; i < vx.Domain.Size(); i++ {
		di = vx.Domain.Value(i)
		vx.Assign(di)
		anyValueSatisfiesArc := false
		if vy.Assigned {
			anyValueSatisfiesArc = a.c.IsSatisfied(assignment)
		} else {
			for j := 0; j < vy.Domain.Size(); j++ {
				dj = vy.Domain.Value(j)
				vy.Assign(dj)
				if a.c.IsSatisfied(assignment) {
					anyValueSatisfiesArc = true
					break
				}
			}
			vy.Unassign()
		}
		if !anyValueSatisfiesArc {
			vx.Domain.Remove(di)
			i--
			revised = true
		}
	}

	vx.Unassign()

	return revised
}

var AC3 csp.InferenceFunc = func(assignment csp.Assignment, constraints []csp.Constraint, varIdx int) (csp.Assignment, bool) {
	as := assignment.Copy()
	//constraints = as.Variables[varIdx].Constraints

	// initially populate the queue with the arcs of the newly assigned variable
	queue := arcQueue{}

	for _, c := range constraints {
		if !c.IsBinaryConstraint() {
			continue
		}
		cIndices := c.AppliesTo()
		if !as.Variable(cIndices[0]).Assigned {
			queue.push(&arc{x: cIndices[0], y: cIndices[1], c: c})
		}
		if !as.Variable(cIndices[1]).Assigned {
			queue.push(&arc{x: cIndices[1], y: cIndices[0], c: c})
		}
	}

	for !queue.empty() {
		a := queue.shift()
		varIdx := a.x
		variable := as.Variable(varIdx)

		if ac3Revise(as, a) {
			if variable.Domain.Size() == 0 {
				return assignment, false
			}

			var c csp.Constraint
			for i := 0; i < len(variable.Constraints); i++ {
				c = variable.Constraints[i]
				if !c.IsBinaryConstraint() {
					continue
				}
				cIndices := c.AppliesTo()
				neighborIdx := cIndices[0]
				if neighborIdx == varIdx {
					neighborIdx = cIndices[1]
				}
				if neighborIdx != a.y && !as.Variable(neighborIdx).Assigned {
					queue.push(&arc{x: neighborIdx, y: varIdx, c: c})
				}
			}
		}
	}

	return as, true
}
