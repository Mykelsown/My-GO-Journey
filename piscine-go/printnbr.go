package student

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	// Handle minimum int value
	// On 64-bit Go: -9223372036854775808
	if n == -9223372036854775808 {
		z01.PrintRune('-')
		// Print "9223372036854775808" by recursion
		PrintNbr(922337203685477580) // print first part
		z01.PrintRune('8')
		return
	}

	// Handle negative values
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	// Recursively print digits
	if n >= 10 {
		PrintNbr(n / 10)
	}

	z01.PrintRune(rune('0' + n%10))
}
