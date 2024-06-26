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

func createVariables(csp CSP) []Variable {
	variables := make([]Variable, len(csp.Domains()))

	for i, d := range csp.Domains() {
		variables[i] = Variable{
			Index:    i,
			Assigned: false,
			Domain:   NewDomain(d),
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
		result[i] = variables[i].Value()
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
