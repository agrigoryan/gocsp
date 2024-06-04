package csp

type Constraint interface {
	AppliesToVariable(index int) bool
	IsSatisfied(solver Solver) bool
}
