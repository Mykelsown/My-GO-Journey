package student

import "github.com/01-edu/z01"

func PrintComb() {
	for a := '0'; a <= '9'; a++ {
		for b := a + 1; b <= '9'; b++ {
			for c := b + 1; c <= '9'; c++ {

				// print the three digits
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(c)

				// check if it's the last combination: 789
				if a == '7' && b == '8' && c == '9' {
					// do not print comma/space
					z01.PrintRune('\n')
				} else {
					// print comma and space
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}
