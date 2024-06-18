package csp

import "slices"

type Assignment struct {
	Variables []Variable
	Domains   []Domain
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

func (a Assignment) Copy() Assignment {
	return Assignment{
		Variables: slices.Clone(a.Variables),
		Domains:   slices.Clone(a.Domains),
	}
}

func (a Assignment) AssignedValue(idx int) (Value, bool) {
	return a.Variables[idx].Value, a.Variables[idx].Assigned
}
