package student

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	comb := make([]int, n)
	backtrack(0, 0, n, comb)
	z01.PrintRune('\n')
}

func backtrack(idx int, start int, n int, comb []int) {
	if idx == n {
		printComb(comb)
		return
	}

	for d := start; d <= 9-(n-idx-1); d++ {
		comb[idx] = d
		backtrack(idx+1, d+1, n, comb)
	}
}

var firstPrinted = false

func printComb(comb []int) {
	if firstPrinted {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	} else {
		firstPrinted = true
	}

	for _, d := range comb {
		z01.PrintRune(rune('0' + d))
	}
}
