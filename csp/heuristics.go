package csp

import "math"

// VariableSelector - a heuristic to pick the next variable to assign
type VariableSelector interface {
	SelectNextVariable(assignment *Assignment) int
}

type VariableSelectorFunc func(assignment *Assignment) int

func (f VariableSelectorFunc) SelectNextVariable(assignment *Assignment) int {
	return f(assignment)
}

var NextUnassignedVariableSelector VariableSelectorFunc = func(assignment *Assignment) int {
	for i := 0; i < assignment.NumVariables(); i++ {
		if !assignment.Variable(i).Assigned {
			return i
		}
	}
	panic("all variables are assigned")
}

// MRVVariableSelector - Minimum Remaining Values heuristic implementation
var MRVVariableSelector VariableSelectorFunc = func(assignment *Assignment) int {
	minDomainSize := math.MaxInt32
	varIdx := -1
	for i := 0; i < assignment.NumVariables(); i++ {
		if assignment.Assigned(i) {
			continue
		}
		if varIdx == -1 || assignment.DomainSize(i) < minDomainSize {
			varIdx = i
			minDomainSize = assignment.DomainSize(varIdx)
		}
	}
	if varIdx == -1 {
		panic("all variables are assigned")
	}
	return varIdx
}

// ValueSelector - a heuristic to pick value to assign to a variable
type ValueSelector interface {
	SelectNextValue(assigment *Assignment, varIndex int) int
}

type ValueSelectorFunc func(assignment *Assignment, varIndex int) int

func (f ValueSelectorFunc) SelectNextValue(assignment *Assignment, varIndex int) int {
	return f(assignment, varIndex)
}

var FirstDomainValueSelector ValueSelectorFunc = func(assignment *Assignment, varIdx int) int {
	if assignment.DomainSize(varIdx) == 0 {
		panic("domain size is 0")
	}
	var idx int
	assignment.RangeDomain(varIdx, func(i int) bool {
		idx = i
		return true
	})
	return idx
}
