package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	var points point
	setPoint(&points)

	output := [14]rune{120, 32, 61, 32, 52, 50, 44, 32, 121, 32, 61, 32, 50, 49}

	for _, val := range output {
		z01.PrintRune(val)
	}

	z01.PrintRune('\n')
}
