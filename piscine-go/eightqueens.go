package student

import "github.com/01-edu/z01"

func EightQueens() {
	var positions [8]int
	solve(0, &positions)
}

func solve(col int, positions *[8]int) {
	if col == 8 {
		printSolution(*positions)
		return
	}

	for row := 1; row <= 8; row++ {
		if isSafe(col, row, *positions) {
			positions[col] = row
			solve(col+1, positions)
		}
	}
}

func isSafe(col, row int, positions [8]int) bool {
	for c := 0; c < col; c++ {
		r := positions[c]
		if r == row || r+col-c == row || r-col+c == row {
			return false
		}
	}
	return true
}

func printSolution(positions [8]int) {
	for _, p := range positions {
		z01.PrintRune(rune('0' + p))
	}
	z01.PrintRune('\n')
}
