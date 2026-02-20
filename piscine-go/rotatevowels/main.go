package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		z01.PrintRune('\n')
		return
	}

	var vowels []rune
	for _, arg := range args {
		for _, r := range arg {
			if isVowel(r) {
				vowels = append(vowels, r)
			}
		}
	}

	if len(vowels) > 1 {
		for i, j := 0, len(vowels)-1; i < j; i, j = i+1, j-1 {
			vowels[i], vowels[j] = vowels[j], vowels[i]
		}
	}

	i := 0
	for a, arg := range args {
		for _, r := range arg {
			if isVowel(r) {
				z01.PrintRune(vowels[i])
				i++
			} else {
				z01.PrintRune(r)
			}
		}
		if a != len(args)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
