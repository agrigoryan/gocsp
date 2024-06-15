package csp

type AssignedValues interface {
	// AssignedValue - return the value of the assigned variable at the given index.
	// If the variable at the given index is not assigned, return
	AssignedValue(idx int) (Value, bool)
}

type ConstraintFunc func(indices []int, assignment AssignedValues) bool

type Constraint struct {
	indices      []int
	checkingFunc ConstraintFunc
}

func (c Constraint) IsSatisfied(assignment AssignedValues) bool {
	return c.checkingFunc(c.indices, assignment)
}

func (c Constraint) AppliesTo() []int {
	return c.indices
}

func (c Constraint) IsBooleanConstraint() bool {
	return len(c.indices) == 2
}

func AllDiffConstraintFunc(indices []int, assignment AssignedValues) bool {
	for i := 0; i < len(indices); i++ {
		val1, ok := assignment.AssignedValue(indices[i])
		if !ok {
			continue
		}
		for j := i + 1; j < len(indices); j++ {
			val2, ok := assignment.AssignedValue(indices[j])
			if ok && val1 == val2 {
				return false
			}
		}
	}
	return true
}

func NewConstraint(indices []int, checkingFunc ConstraintFunc) Constraint {
	return Constraint{
		indices:      indices,
		checkingFunc: checkingFunc,
	}
}

func NewBinaryConstraint(idx1, idx2 int, checkingFunc ConstraintFunc) Constraint {
	return NewConstraint([]int{idx1, idx2}, checkingFunc)
}

func NewAllDiffConstraint(indices []int) Constraint {
	return NewConstraint(indices, AllDiffConstraintFunc)
}
