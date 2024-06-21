package csp

import "math"

// VariableSelector - a strategy to pick the next variable to assign
type VariableSelector interface {
	SelectNextVariable(assignment Assignment) int
}

type VariableSelectorFunc func(assignment Assignment) int

func (f VariableSelectorFunc) SelectNextVariable(assignment Assignment) int {
	return f(assignment)
}

var NextUnassignedVariableSelector VariableSelectorFunc = func(assignment Assignment) int {
	for i := 0; i < assignment.NumVariables(); i++ {
		if !assignment.Variable(i).Assigned {
			return i
		}
	}
	panic("all variables are assigned")
}

// MRVVariableSelector - Minimum Remaining Values heuristic implementation
var MRVVariableSelector VariableSelectorFunc = func(assignment Assignment) int {
	minDomainSize := math.MaxInt32
	varIdx := -1
	for i := 0; i < assignment.NumVariables(); i++ {
		v := assignment.Variable(i)
		if !v.Assigned && v.Domain.Size() < minDomainSize {
			minDomainSize = v.Domain.Size()
			varIdx = i
		}
	}
	if varIdx == -1 {
		panic("all variables are assigned")
	}
	return varIdx
}

// ValueSelector - a strategy to pick value to assign to a variable
type ValueSelector interface {
	SelectVariableValue(assigment Assignment, varIndex int) Value
}

type ValueSelectorFunc func(assignment Assignment, varIndex int) Value

func (f ValueSelectorFunc) SelectVariableValue(assignment Assignment, varIndex int) Value {
	return f(assignment, varIndex)
}

var FirstDomainValueSelector ValueSelectorFunc = func(assignment Assignment, varIndex int) Value {
	return assignment.Variable(varIndex).Domain.Value(0)
}

// Solver - generic interface to solve CSP
type Solver interface {
	Solve(csp CSP) []Value
}

type SimpleSolver struct {
	variableSelector VariableSelector
	valueSelector    ValueSelector
	inference        Inference
}

func (s *SimpleSolver) Solve(csp CSP) []Value {
	assignment := Assignment{
		Variables: createVariables(csp),
	}
	res := s.solveAssignment(assignment, csp.Constraints())
	return res
}

func (s *SimpleSolver) solveAssignment(assignment Assignment, constraints []Constraint) []Value {
	if assignment.IsComplete(constraints) {
		return variableValues(assignment.Variables)
	}

	varIdx := s.variableSelector.SelectNextVariable(assignment)
	origDomain := assignment.Variable(varIdx).Domain.Copy()

	for assignment.Variable(varIdx).Domain.Size() > 0 {
		value := s.valueSelector.SelectVariableValue(assignment, varIdx)
		variable := assignment.Variable(varIdx)
		variable.Assign(value)
		if !assignment.IsConsistent(variable.Constraints) {
			variable.Domain.Remove(value)
			continue
		}

		updatedAssignment, ok := s.inference.Inference(assignment, constraints, varIdx)
		if ok {
			assignment = updatedAssignment
			res := s.solveAssignment(assignment, constraints)
			if res != nil {
				assignment.Variable(varIdx).Domain.RemoveAllBut(value)
				return res
			}
		}

		assignment.Variable(varIdx).Domain.Remove(value)
	}

	variable := assignment.Variable(varIdx)
	variable.Unassign()
	variable.Domain = origDomain

	return nil
}

func createVariables(csp CSP) []Variable {
	variables := make([]Variable, len(csp.Domains()))

	for i, d := range csp.Domains() {
		variables[i] = Variable{
			Index:    i,
			Assigned: false,
			Domain:   d,
		}
	}

	for _, c := range csp.Constraints() {
		for _, i := range c.AppliesTo() {
			variables[i].Constraints = append(variables[i].Constraints, c)
		}
	}

	return variables
}

func variableValues(variables []Variable) []Value {
	result := make([]Value, len(variables))
	for i := range variables {
		result[i] = variables[i].Value
	}
	return result
}

func NewSimpleSolver(variableSelector VariableSelector, valueSelector ValueSelector, inference Inference) *SimpleSolver {
	return &SimpleSolver{
		variableSelector: variableSelector,
		valueSelector:    valueSelector,
		inference:        inference,
	}
}
