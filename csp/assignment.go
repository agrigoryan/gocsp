package csp

import "slices"

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

func (a Assignment) Domain(idx int) Domain {
	return a.Variables[idx].Domain
}

func (a Assignment) SetDomain(idx int, domain Domain) {
	a.Variables[idx].Domain = domain
}

func (a Assignment) AssignedValue(idx int) (Value, bool) {
	return a.Variables[idx].Value, a.Variables[idx].Assigned
}

func (a Assignment) Copy() Assignment {
	return Assignment{
		Variables: slices.Clone(a.Variables),
	}
}
