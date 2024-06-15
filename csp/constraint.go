package csp

type ConstraintFunc func(indices []int, assignment Assignment) bool

type Constraint struct {
	indices      []int
	checkingFunc ConstraintFunc
}

func (c Constraint) IsSatisfied(assignment Assignment) bool {
	return c.checkingFunc(c.indices, assignment)
}

func (c Constraint) AppliesTo() []int {
	return c.indices
}

func (c Constraint) IsBooleanConstraint() bool {
	return len(c.indices) == 2
}

func AllDiffConstraintFunc(indices []int, assignment Assignment) bool {
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

func BinaryAllDiffConstraintFunc(indices []int, assignment Assignment) bool {
	return !assignment.Variables[indices[0]].Assigned ||
		!assignment.Variables[indices[1]].Assigned ||
		assignment.Variables[indices[0]].Value != assignment.Variables[indices[1]].Value
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

func NewBinaryAllDiffConstraint(idx1, idx2 int) Constraint {
	return NewBinaryConstraint(idx1, idx2, BinaryAllDiffConstraintFunc)
}
