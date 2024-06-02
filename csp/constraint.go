package csp

type Constraint interface {
	IsSatisfied(solver Solver) bool
}
