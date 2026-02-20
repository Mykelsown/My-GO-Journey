package student

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	i := '0'

	for i <= '9' {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for l := '0'; l <= '9'; l++ {
					if i > k || (i == k && j >= l) {
						continue
					}
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					z01.PrintRune(l)

					if i == '9' && j == '8' && k == '9' && l == '9' {
						z01.PrintRune('\n')
					} else {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
		i = i + 1
	}
}
