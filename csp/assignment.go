package csp

import "slices"

type Assignment struct {
	Variables []Variable
	Domains   []Domain
}

func (a Assignment) IsConsistent(constraints []Constraint) bool {
	for _, c := range constraints {
		if !c.IsSatisfied(a) {
			return false
		}
	}
	return true
}

func (a Assignment) IsComplete(constraints []Constraint) bool {
	for _, v := range a.Variables {
		if !v.Assigned {
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
