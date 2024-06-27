package csp

import "fmt"

type SolverProgressListener interface {
	OnSolverProgress(assignment *Assignment)
}

type BacktrackingSolver struct {
	variableSelector VariableSelector
	valueSelector    ValueSelector
	inference        Inference

	stepsCounter int

	Listener SolverProgressListener
}

func (s *BacktrackingSolver) Solve(csp CSP) []Value {
	assignment := createInitialAssignment(csp)
	s.stepsCounter = 0
	res := s.solveAssignment(assignment, csp.Constraints())
	fmt.Printf("solved in %d steps\n", s.stepsCounter)
	return res
}

func (s *BacktrackingSolver) solveAssignment(assignment *Assignment, constraints []Constraint) []Value {
	s.stepsCounter++
	if s.Listener != nil {
		s.Listener.OnSolverProgress(assignment)
	}

	if assignment.IsComplete(constraints) {
		return variableValues(assignment)
	}

	varIdx := s.variableSelector.SelectNextVariable(assignment)
	origDomain := assignment.Domain(varIdx).Clone()

	for assignment.DomainSize(varIdx) > 0 {
		valueIdx := s.valueSelector.SelectNextValue(assignment, varIdx)
		assignment.Assign(varIdx, valueIdx)
		if !assignment.IsConsistent(assignment.Constraints(varIdx)) {
			assignment.Unset(varIdx, valueIdx)
			continue
		}

		nextAssignment := assignment

		if s.inference != nil {
			var ok bool
			nextAssignment, ok = s.inference.Inference(nextAssignment.Clone(), constraints, varIdx)
			if !ok {
				assignment.Unset(varIdx, valueIdx)
				continue
			}
		}

		if res := s.solveAssignment(nextAssignment, constraints); res != nil {
			return res
		}

		assignment.Unset(varIdx, valueIdx)
	}

	assignment.Unassign(varIdx)
	assignment.OverwriteDomain(varIdx, origDomain)

	return nil
}

func createInitialAssignment(csp CSP) *Assignment {
	numDomains := len(csp.Domains())
	variables := make([]Variable, numDomains)
	domains := make([]Domain, numDomains)

	for i, d := range csp.Domains() {
		variables[i] = Variable{
			Index:    i,
			Assigned: false,
		}
		domains[i] = NewDomain(d)
	}

	for _, c := range csp.Constraints() {
		for _, i := range c.AppliesTo() {
			variables[i].Constraints = append(variables[i].Constraints, c)
		}
	}

	return &Assignment{
		Variables: variables,
		Domains:   domains,
	}
}

func variableValues(assignment *Assignment) []Value {
	result := make([]Value, assignment.NumVariables())
	for i := range result {
		result[i], _ = assignment.AssignedValue(i)
	}
	return result
}

func NewBacktrackingSolver(variableSelector VariableSelector, valueSelector ValueSelector, inference Inference) *BacktrackingSolver {
	return &BacktrackingSolver{
		variableSelector: variableSelector,
		valueSelector:    valueSelector,
		inference:        inference,
	}
}
