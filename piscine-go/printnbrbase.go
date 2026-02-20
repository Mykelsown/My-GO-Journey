package student

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
	}

	printBaseRec(nbr, base, len(base))
}

func printBaseRec(n int, base string, size int) {
	if n >= size || n <= -size {
		printBaseRec(n/size, base, size)
	}

	idx := n % size
	if idx < 0 {
		idx = -idx
	}

	z01.PrintRune(rune(base[idx]))
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' || seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}
