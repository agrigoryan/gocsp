package csp

import "slices"

type Assignment struct {
	Variables []Variable
	Domains   []Domain
}

func (a *Assignment) IsConsistent(constraints []Constraint) bool {
	for i := 0; i < len(constraints); i++ {
		if !constraints[i].IsSatisfied(a) {
			return false
		}
	}
	return true
}

func (a *Assignment) IsComplete(constraints []Constraint) bool {
	for i := 0; i < len(a.Variables); i++ {
		if !a.Variables[i].Assigned {
			return false
		}
	}
	return a.IsConsistent(constraints)
}

func (a *Assignment) NumVariables() int {
	return len(a.Variables)
}

func (a *Assignment) Variable(idx int) *Variable {
	return &a.Variables[idx]
}

func (a *Assignment) Constraints(varIdx int) []Constraint {
	return a.Variables[varIdx].Constraints
}

func (a *Assignment) Domain(varIdx int) *Domain {
	return &a.Domains[varIdx]
}

func (a *Assignment) DomainValue(varIdx int, valIdx int) Value {
	return a.Domains[varIdx].Value(valIdx)
}

func (a *Assignment) OverwriteDomain(varIdx int, domain Domain) {
	a.Domains[varIdx] = domain
}

func (a *Assignment) Assigned(varIdx int) bool {
	return a.Variables[varIdx].Assigned
}

func (a *Assignment) Assign(varIdx, valIdx int) {
	a.Variables[varIdx].Assign(valIdx)
}

func (a *Assignment) Unassign(varIdx int) {
	a.Variables[varIdx].Unassign()
}

func (a *Assignment) DomainSize(varIdx int) int {
	return a.Domains[varIdx].Size()
}

func (a *Assignment) Set(varIdx, valIdx int) {
	a.Domains[varIdx].Set(valIdx)
}

func (a *Assignment) Unset(varIdx, valIdx int) {
	a.Domains[varIdx].Unset(valIdx)
}

func (a *Assignment) Contains(varIdx, valIdx int) {
	a.Domains[varIdx].Contains(valIdx)
}

func (a *Assignment) RangeDomain(varIdx int, fn func(int) bool) {
	a.Domains[varIdx].Range(fn)
}

func (a *Assignment) FilterDomain(varIdx int, fn func(int) bool) {
	a.Domains[varIdx].Filter(fn)
}

func (a *Assignment) AssignedValueIdx(idx int) (int, bool) {
	return a.Variables[idx].ValueIdx, a.Variables[idx].Assigned
}

func (a *Assignment) AssignedValue(idx int) (Value, bool) {
	if !a.Variables[idx].Assigned {
		return 0, false
	}
	return a.DomainValue(idx, a.Variables[idx].ValueIdx), true
}

func (a *Assignment) Clone() *Assignment {
	clonedDomains := slices.Clone(a.Domains)
	for i := range clonedDomains {
		clonedDomains[i] = clonedDomains[i].Clone()
	}
	return &Assignment{
		Variables: a.Variables,
		Domains:   clonedDomains,
	}
}
