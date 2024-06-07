package csp

// VariablePicker - a strategy to pick the next variable to assign
type VariablePicker interface {
	NextVariableIndex(assignment Assignment) int
}

type VariablePickerFunc func(assignment Assignment) int

func (f VariablePickerFunc) NextVariableIndex(assignment Assignment) int {
	return f(assignment)
}

var NextUnassignedVariablePicker VariablePickerFunc = func(assignment Assignment) int {
	for i, v := range assignment.Variables {
		if !v.Assigned {
			return i
		}
	}
	panic("all variables are assigned")
}

// ValuePicker - a strategy to pick value to assign to a variable
type ValuePicker interface {
	VariableValue(assigment Assignment, varIndex int) Value
}

type ValuePickerFunc func(assignment Assignment, varIndex int) Value

func (f ValuePickerFunc) VariableValue(assignment Assignment, varIndex int) Value {
	return f(assignment, varIndex)
}

var FirstValidValuePicker ValuePickerFunc = func(assignment Assignment, varIndex int) Value {
	return assignment.Domains[varIndex].Values()[0]
}

// Solver - generic interface to solve CSP
type Solver interface {
	Solve(csp CSP) []Value
}

type SimpleSolver struct {
	variablePicker VariablePicker
	valuePicker    ValuePicker
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

	newAssignment := assignment.Copy()
	varIdx := s.variablePicker.NextVariableIndex(assignment)
	domain := newAssignment.Domains[varIdx]
	for {
		value := s.valuePicker.VariableValue(newAssignment, varIdx)
		assignedVar := newAssignment.Variables[varIdx].Assign(value)
		newAssignment.Variables[varIdx] = assignedVar
		if !newAssignment.IsConsistent(assignedVar.Constraints) {
			domain.Remove(value)
			if domain.Size() == 0 {
				return nil
			}
		} else {
			break
		}
	}

	return s.solveAssignment(newAssignment, constraints)
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

func NewSimpleSolver(variablePicker VariablePicker, valuePicker ValuePicker) *SimpleSolver {
	return &SimpleSolver{
		variablePicker: variablePicker,
		valuePicker:    valuePicker,
	}
}
