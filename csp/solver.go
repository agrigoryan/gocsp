package csp

type Solver struct {
	variables   []Variable
	constraints []Constraint
}

func (s *Solver) Variables() []Variable {
	return s.variables
}

func NewSolver() {

}
