## GoCSP

A simple Go library for solving constraint satisfaction problems using backtracking.

Comes with some basic algorithms - like AC3, Forward Checking, MRV for heuristics.

### Overview

NQueens solving:
```go
problem := nqueens.New(20)
solver := csp.NewBacktrackingSolver(csp.NextUnassignedVariableSelector, csp.FirstDomainValueSelector, inference.FwdCheck)
result := solver.Solve(problem)
fmt.Println(result)
```

Map coloring:
```go
states := []string{"WA", "NT", "SA", "Q", "NSW", "V", "T"}
domains := make([]csp.ValueSet, 0, len(states))
for i := range states {
    domains = append(domains, csp.ValueSet{1, 2, 3})
}
constraints := []csp.Constraint{
    csp.NewAllDiffConstraint([]int{0, 1}),
    csp.NewAllDiffConstraint([]int{0, 2}),
    csp.NewAllDiffConstraint([]int{1, 2}),
    csp.NewAllDiffConstraint([]int{1, 3}),
    csp.NewAllDiffConstraint([]int{2, 3}),
    csp.NewAllDiffConstraint([]int{2, 4}),
    csp.NewAllDiffConstraint([]int{2, 5}),
    csp.NewAllDiffConstraint([]int{3, 4}),
    csp.NewAllDiffConstraint([]int{4, 5}),
}

problem := csp.NewGenericCSP(domains, constraints)
solver := csp.NewBacktrackingSolver(csp.MRVVariableSelector,  csp.FirstDomainValueSelector, nil)
result := solver.Solve(problem)
fmt.Println(result)
```

Sudoku solver example: github.com/agrigoryan/sudoku
