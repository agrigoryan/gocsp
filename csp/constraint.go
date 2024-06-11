package csp

type Constraint interface {
	IsSatisfied(assignment Assignment) bool

	// AppliesTo - The list of indices this constraint applies to
	AppliesTo() []int
}

type ConstraintFunc func(indices []int, assignment Assignment) bool

type GenericConstraint struct {
	indices      []int
	checkingFunc ConstraintFunc
}

func (c GenericConstraint) IsSatisfied(assignment Assignment) bool {
	return c.checkingFunc(c.indices, assignment)
}

func (c GenericConstraint) AppliesTo() []int {
	return c.indices
}

func (c GenericConstraint) IsBooleanConstraint() bool {
	return len(c.indices) == 2
}

func AllDifferentConstraintFunc(indices []int, assignment Assignment) bool {
	for i := 0; i < len(indices); i++ {
		var1 := assignment.Variables[indices[i]]
		if !var1.Assigned {
			continue
		}
		for j := i + 1; j < len(indices); j++ {
			var2 := assignment.Variables[indices[j]]
			if var2.Assigned && var1.Value == var2.Value {
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
	return GenericConstraint{
		indices:      []int{idx1, idx2},
		checkingFunc: checkingFunc,
	}
}

func NewAllDifferentConstraint(indices []int) GenericConstraint {
	return NewConstraint(indices, AllDifferentConstraintFunc)
}
