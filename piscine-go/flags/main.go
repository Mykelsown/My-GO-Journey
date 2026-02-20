package main

import "os"

func printHelp() {
	os.Stdout.WriteString("--insert\n")
	os.Stdout.WriteString("  -i\n")
	os.Stdout.WriteString("\t This flag inserts the string into the string passed as argument.\n")
	os.Stdout.WriteString("--order\n")
	os.Stdout.WriteString("  -o\n")
	os.Stdout.WriteString("\t This flag will behave like a boolean, if it is called it will order the argument.\n")
}

func sortString(s string) string {
	b := []byte(s)
	for i := 0; i < len(b)-1; i++ {
		for j := i + 1; j < len(b); j++ {
			if b[i] > b[j] {
				b[i], b[j] = b[j], b[i]
			}
		}
	}
	return string(b)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printHelp()
		return
	}

	insert := ""
	order := false
	mainStr := ""

	for _, arg := range args {
		if len(arg) >= 9 && arg[:9] == "--insert=" {
			insert = arg[9:]
		} else if len(arg) >= 3 && arg[:3] == "-i=" {
			insert = arg[3:]
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else if arg == "--help" || arg == "-h" {
			printHelp()
			return
		} else {
			mainStr = arg
		}
	}

	if insert != "" {
		mainStr += insert
	}

	if order {
		mainStr = sortString(mainStr)
	}

	os.Stdout.WriteString(mainStr + "\n")
}
