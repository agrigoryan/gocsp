package csp

// VariableSelector - a strategy to pick the next variable to assign
type VariableSelector interface {
	SelectNextVariable(assignment Assignment) int
}

type VariableSelectorFunc func(assignment Assignment) int

func (f VariableSelectorFunc) SelectNextVariable(assignment Assignment) int {
	return f(assignment)
}

var NextUnassignedVariableSelector VariableSelectorFunc = func(assignment Assignment) int {
	for i, v := range assignment.Variables {
		if !v.Assigned {
			return i
		}
	}
	panic("all variables are assigned")
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
	return assignment.Domains[varIndex].Values()[0]
}

// Solver - generic interface to solve CSP
type Solver interface {
	Solve(csp CSP) []Value
}

type SimpleSolver struct {
	variableSelector VariableSelector
	valueSelector    ValueSelector
}

func (s *SimpleSolver) Solve(csp CSP) []Value {
	assignment := Assignment{
		Variables: createVariables(csp),
		Domains:   csp.Domains(),
	}
	return s.solveAssignment(assignment, csp.Constraints())
}

func (s *SimpleSolver) solveAssignment(assignment Assignment, constraints []Constraint) []Value {
	if assignment.IsComplete(constraints) {
		return variableValues(assignment.Variables)
	}

	varIdx := s.variableSelector.SelectNextVariable(assignment)
	domain := assignment.Domains[varIdx]
	origDomain := domain.ShallowCopy()

	for domain.Size() > 0 {
		value := s.valueSelector.SelectVariableValue(assignment, varIdx)
		assignedVar := assignment.Variables[varIdx].Assign(value)
		assignment.Variables[varIdx] = assignedVar
		if !assignment.IsConsistent(assignedVar.Constraints) {
			domain.Remove(value)
			continue
		}
		// TODO: run forward checking or other consistency checks here on the new assignment
		res := s.solveAssignment(assignment, constraints)
		if res != nil {
			return res
		} else {
			domain.Remove(value)
		}
	}

	assignment.Variables[varIdx] = assignment.Variables[varIdx].Unassign()
	assignment.Domains[varIdx] = origDomain

	return nil
}

func createVariables(csp CSP) []Variable {
	variables := make([]Variable, len(csp.Domains()))

	for i := range csp.Domains() {
		variables[i] = Variable{
			Index:    i,
			Assigned: false,
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

func NewSimpleSolver(variableSelector VariableSelector, valueSelector ValueSelector) *SimpleSolver {
	return &SimpleSolver{
		variableSelector: variableSelector,
		valueSelector:    valueSelector,
	}
}
