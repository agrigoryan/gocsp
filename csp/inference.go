package csp

type Inference interface {
	Inference(assignment *Assignment, constraints []Constraint, varIdx int) bool
}

type InferenceFunc func(assignment *Assignment, constraints []Constraint, varIdx int) bool

func (f InferenceFunc) Inference(assignment *Assignment, constraints []Constraint, varIdx int) bool {
	return f(assignment, constraints, varIdx)
}

type CompositeInference struct {
	inferences []Inference
}

func (ci *CompositeInference) Inference(assignment *Assignment, constraints []Constraint, varIdx int) bool {
	for _, inf := range ci.inferences {
		if !inf.Inference(assignment, constraints, varIdx) {
			return false
		}
	}
	return true
}

func NewCompositeInference(inferences ...Inference) CompositeInference {
	return CompositeInference{
		inferences: inferences,
	}
}
