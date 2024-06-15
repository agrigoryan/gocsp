package csp

type AssignedValues interface {
	// AssignedValue - return the value of the assigned variable at the given index.
	// If the variable at the given index is not assigned, return
	AssignedValue(idx int) (Value, bool)
}

type Constraint interface {
	IsSatisfied(assignment AssignedValues) bool

	// AppliesTo - The list of indices this constraint applies to
	AppliesTo() []int
}

type ConstraintFunc func(indices []int, assignment AssignedValues) bool

type GenericConstraint struct {
	indices      []int
	checkingFunc ConstraintFunc
}

func (c GenericConstraint) IsSatisfied(assignment AssignedValues) bool {
	return c.checkingFunc(c.indices, assignment)
}

func (c GenericConstraint) AppliesTo() []int {
	return c.indices
}

func (c GenericConstraint) IsBooleanConstraint() bool {
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

func NewConstraint(indices []int, checkingFunc ConstraintFunc) GenericConstraint {
	return GenericConstraint{
		indices:      indices,
		checkingFunc: checkingFunc,
	}
}

func NewBinaryConstraint(idx1, idx2 int, checkingFunc ConstraintFunc) GenericConstraint {
	return NewConstraint([]int{idx1, idx2}, checkingFunc)
}

func NewAllDiffConstraint(indices []int) GenericConstraint {
	return NewConstraint(indices, AllDiffConstraintFunc)
}
