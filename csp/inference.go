package csp

type Inference interface {
	Inference(assignment Assignment, constraints []Constraint, varIdx int) (Assignment, bool)
}

type InferenceFunc func(assignment Assignment, constraints []Constraint, varIdx int) (Assignment, bool)

func (f InferenceFunc) Inference(assignment Assignment, constraints []Constraint, varIdx int) (Assignment, bool) {
	return f(assignment, constraints, varIdx)
}

var NoInferenceFunc InferenceFunc = func(assignment Assignment, constraints []Constraint, varIdx int) (Assignment, bool) {
	return assignment, true
}

var ForwardCheckInferenceFunc InferenceFunc = func(assignment Assignment, constraints []Constraint, varIdx int) (Assignment, bool) {
	panic("not implemented")
}
