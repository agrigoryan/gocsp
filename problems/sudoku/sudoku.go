package sudoku

import "github.com/agrigoryan/gocsp/csp"

func sudokuConstraints() []csp.Constraint {
	constraints := make([]csp.Constraint, 0, 3*9)

	for i := 0; i < 9; i++ {
		indices := make([]int, 9)
		for j := 0; j < 9; j++ {
			indices[j] = i*9 + j
		}
		constraints = append(constraints, csp.NewAllDiffConstraint(indices))

		indices = make([]int, 9)
		for j := 0; j < 9; j++ {
			indices[j] = 9*j + i
		}
		constraints = append(constraints, csp.NewAllDiffConstraint(indices))

		indices = make([]int, 9)
		blockX := i % 3
		blockY := i / 3
		for j := 0; j < 9; j++ {
			cx := j % 3
			cy := j / 3
			indices[j] = (blockY*3+cy)*9 + cx + blockX*3
		}
		constraints = append(constraints, csp.NewAllDiffConstraint(indices))
	}

	return constraints
}

func New(board Board) *csp.GenericCSP {
	domains := make([]csp.ValueSet, len(board))
	for i := 0; i < 81; i++ {
		if board[i] == 0 {
			domains[i] = csp.ValueSet{1, 2, 3, 4, 5, 6, 7, 8, 9}
		} else {
			domains[i] = csp.ValueSet{csp.Value(board[i])}
		}
	}

	return csp.NewGenericCSP(domains, sudokuConstraints())
}
