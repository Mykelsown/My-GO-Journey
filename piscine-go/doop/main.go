package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	aStr := os.Args[1]
	op := os.Args[2]
	bStr := os.Args[3]

	a, okA := Atoi(aStr)
	b, okB := Atoi(bStr)

	if !okA || !okB {
		return
	}

	switch op {
	case "+":
		res, ok := Add(a, b)
		if !ok {
			return
		}
		PrintInt(res)
		os.Stdout.WriteString("\n")

	case "-":
		res, ok := Sub(a, b)
		if !ok {
			return
		}
		PrintInt(res)
		os.Stdout.WriteString("\n")

	case "*":
		res, ok := Mul(a, b)
		if !ok {
			return
		}
		PrintInt(res)
		os.Stdout.WriteString("\n")

	case "/":
		if b == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		PrintInt(a / b)
		os.Stdout.WriteString("\n")

	case "%":
		if b == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		PrintInt(a % b)
		os.Stdout.WriteString("\n")

	default:
		return
	}
}

//
// ================= ATOI =================
//   Manual conversion without strconv
//

func Atoi(s string) (int64, bool) {
	if len(s) == 0 {
		return 0, false
	}

	sign := int64(1)
	i := 0

	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	var n int64 = 0
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		d := int64(s[i] - '0')

		// overflow check
		if n > (9223372036854775807-d)/10 {
			return 0, false
		}

		n = n*10 + d
	}

	return n * sign, true
}

//
// ========= SAFE ARITHMETIC WITH OVERFLOW =========
//

func Add(a, b int64) (int64, bool) {
	if (b > 0 && a > 9223372036854775807-b) ||
		(b < 0 && a < -9223372036854775808-b) {
		return 0, false
	}
	return a + b, true
}

func Sub(a, b int64) (int64, bool) {
	return Add(a, -b)
}

func Mul(a, b int64) (int64, bool) {
	if a == 0 || b == 0 {
		return 0, true
	}

	// overflow checks
	if a > 0 && b > 0 && a > 9223372036854775807/b {
		return 0, false
	}
	if a > 0 && b < 0 && b < -9223372036854775808/a {
		return 0, false
	}
	if a < 0 && b > 0 && a < -9223372036854775808/b {
		return 0, false
	}
	if a < 0 && b < 0 && a < 9223372036854775807/b {
		return 0, false
	}

	return a * b, true
}

//
// ================ PrintInt =================
//

func PrintInt(n int64) {
	if n == 0 {
		os.Stdout.Write([]byte("0"))
		return
	}

	if n < 0 {
		os.Stdout.Write([]byte("-"))
		n = -n
	}

	var buf [20]byte
	i := 0

	for n > 0 {
		buf[i] = byte(n%10) + '0'
		n /= 10
		i++
	}

	for j := i - 1; j >= 0; j-- {
		os.Stdout.Write([]byte{buf[j]})
	}
}
