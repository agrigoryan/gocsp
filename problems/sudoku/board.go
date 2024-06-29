package sudoku

import (
	"strings"
)

type Board []byte

func (b Board) IsValid() bool {
	if len(b) != 81 {
		return false
	}
	for i := 0; i < 81; i++ {
		if b[i] < 0 || b[i] > 9 {
			return false
		}
	}
	return true
}

func (b Board) IsSolved() bool {
	if !b.IsValid() {
	}
	for i := 0; i < 81; i++ {
		if b[i] == 0 {
			return false
		}
	}
	return true
}

func NewBoard() Board {
	return make(Board, 81)
}

func NewBoardFromString(input string) Board {
	board := NewBoard()
	for i := range board {
		board[i] = 0xFF
	}
	idx := 0
	for _, r := range []rune(input) {
		if r == '0' || r == '.' {
			board[idx] = 0
			idx++
		} else if r >= '1' && r <= '9' {
			board[idx] = byte(r - '0')
			idx++
		}
	}
	return board
}

func (b Board) String() string {
	res := strings.Builder{}
	for i := 0; i < 81; i++ {
		if i%27 == 0 {
			if i > 0 {
				res.WriteString("|\n")
			}
			res.WriteString("*---*---*---*\n")
		} else if i%9 == 0 {
			res.WriteString("|\n")
		}
		if i%3 == 0 {
			res.WriteByte('|')
		}
		if b[i] == 0 {
			res.WriteByte('.')
		} else {
			res.WriteByte('0' + b[i])
		}
	}
	res.WriteString("|\n*---*---*---*\n")
	return res.String()
}
