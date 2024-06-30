## GoCSP

A simple Go library for solving constraint satisfaction problems using backtracking.

Comes with some basic algorithms - like AC3, Forward Checking, MVR for heuristics.

### Overview

```go
problem := nqueens.New(20)
solver := csp.NewBacktrackingSolver(csp.NextUnassignedVariableSelector, csp.FirstDomainValueSelector, inference.FwdCheck)
result := solver.Solve(problem)
fmt.Println(result)
```
