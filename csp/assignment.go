package csp

type Assignment struct {
	Variables []Variable
}

func (a Assignment) IsConsistent(constraints []Constraint) bool {
	for i := 0; i < len(constraints); i++ {
		if !constraints[i].IsSatisfied(a) {
			return false
		}
	}
	return true
}

func (a Assignment) IsComplete(constraints []Constraint) bool {
	for i := 0; i < len(a.Variables); i++ {
		if !a.Variables[i].Assigned {
			return false
		}
	}
	return a.IsConsistent(constraints)
}

func (a Assignment) NumVariables() int {
	return len(a.Variables)
}

func (a Assignment) Variable(idx int) *Variable {
	return &a.Variables[idx]
}

func (a Assignment) AssignedValueIdx(idx int) (int, bool) {
	return a.Variables[idx].ValueIdx, a.Variables[idx].Assigned
}

func (a Assignment) Clone() Assignment {
	clonedVars := make([]Variable, len(a.Variables))
	for i, v := range a.Variables {
		v.Copy(&clonedVars[i])
	}
	return Assignment{
		Variables: clonedVars,
	}
}
