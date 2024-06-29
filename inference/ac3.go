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

func ac3Revise(assignment *csp.Assignment, a *arc) bool {
	revised := false

	assignment.RangeDomain(a.x, func(i int) bool {
		assignment.Assign(a.x, i)
		anyValueSatisfiesArc := false
		if assignment.Assigned(a.y) {
			anyValueSatisfiesArc = a.c.IsSatisfied(assignment)
		} else {
			assignment.RangeDomain(a.y, func(j int) bool {
				assignment.Assign(a.y, j)
				if a.c.IsSatisfied(assignment) {
					anyValueSatisfiesArc = true
					return true
				}
				return false
			})
			assignment.Unassign(a.y)
		}
		if !anyValueSatisfiesArc {
			assignment.Unset(a.x, i)
			revised = true
		}
		return false
	})

	assignment.Unassign(a.x)

	return revised
}

var AC3 csp.InferenceFunc = func(assignment *csp.Assignment, constraints []csp.Constraint, varIdx int) bool {
	// optionally limit the initial set of arcs to the constraints of the newly assigned variable
	constraints = assignment.Constraints(varIdx)

	// initially populate the queue with the arcs of the newly assigned variable
	queue := arcQueue{}

	for _, c := range constraints {
		if !c.IsBinaryConstraint() {
			continue
		}
		cIndices := c.AppliesTo()
		if !assignment.Assigned(cIndices[0]) {
			queue.push(&arc{x: cIndices[0], y: cIndices[1], c: c})
		}
		if !assignment.Assigned(cIndices[1]) {
			queue.push(&arc{x: cIndices[1], y: cIndices[0], c: c})
		}
	}

	for !queue.empty() {
		a := queue.shift()
		varIdx := a.x

		if ac3Revise(assignment, a) {
			if assignment.DomainSize(varIdx) == 0 {
				return false
			}

			for _, c := range assignment.Constraints(varIdx) {
				if !c.IsBinaryConstraint() {
					continue
				}
				cIndices := c.AppliesTo()
				neighborIdx := cIndices[0]
				if neighborIdx == varIdx {
					neighborIdx = cIndices[1]
				}
				if neighborIdx != a.y && !assignment.Assigned(neighborIdx) {
					queue.push(&arc{x: neighborIdx, y: varIdx, c: c})
				}
			}
		}
	}

	return true
}
