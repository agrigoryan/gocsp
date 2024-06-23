package csp

type SolverProgressListener interface {
	OnSolverProgress(assignment Assignment)
}

type BacktrackingSolver struct {
	variableSelector VariableSelector
	valueSelector    ValueSelector
	inference        Inference

	Listener SolverProgressListener
}

func (s *BacktrackingSolver) Solve(csp CSP) []Value {
	assignment := Assignment{
		Variables: createVariables(csp),
	}
	res := s.solveAssignment(assignment, csp.Constraints())
	return res
}

func (s *BacktrackingSolver) solveAssignment(assignment Assignment, constraints []Constraint) []Value {
	if s.Listener != nil {
		s.Listener.OnSolverProgress(assignment)
	}

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

		inferredAssignment, ok := s.inference.Inference(assignment, constraints, varIdx)
		if !ok {
			variable.Domain.Remove(value)
			continue
		}

		if res := s.solveAssignment(inferredAssignment, constraints); res != nil {
			return res
		}

		variable.Domain.Remove(value)
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

func NewBacktrackingSolver(variableSelector VariableSelector, valueSelector ValueSelector, inference Inference) *BacktrackingSolver {
	return &BacktrackingSolver{
		variableSelector: variableSelector,
		valueSelector:    valueSelector,
		inference:        inference,
	}
}
