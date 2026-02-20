package main

import "github.com/01-edu/z01"

func main() {
	for i := 97; i < 123; i++ {
		z01.PrintRune(rune(i)) // Using PrintRune per standard Go convention (case-sensitive)
	}
	z01.PrintRune('\n')
}
