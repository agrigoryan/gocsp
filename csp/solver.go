package csp

import (
	"fmt"
)

type SolverResult []Value

type Solver interface {
	Solve(csp CSP) []SolverResult
}

type SolverFunc func(csp CSP) []SolverResult

func (f SolverFunc) Solve(csp CSP) []SolverResult {
	return f(csp)
}

type SimpleSolver struct{}

func (s *SimpleSolver) Solve(csp CSP) []SolverResult {
	domains := csp.Domains()
	variables := make([]Variable, len(domains))
	for i := range domains {
		variables = append(variables, Variable{
			Id:       i,
			Assigned: false,
		})
	}
	// TODO: solve the problem
	return nil
}

func pickNextVariableToAssign(variables []Variable) (int, error) {
	for i, v := range variables {
		if !v.Assigned {
			return i, nil
		}
	}
	return 0, fmt.Errorf("all variables are assigned")
}

func allAssigned(variables []Variable) bool {
	for _, v := range variables {
		if !v.Assigned {
			return false
		}
	}
	return true
}

func NewSimpleSolver() *SimpleSolver {
	return &SimpleSolver{}
}
