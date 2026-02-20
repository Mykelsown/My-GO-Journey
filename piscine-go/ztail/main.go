package main

import (
	"fmt"
	"os"
)

func parsePositiveInt(s string) (int, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("invalid number")
	}
	n := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, fmt.Errorf("invalid number")
		}
		n = n*10 + int(c-'0')
	}
	return n, nil
}

func tailFromOpen(f *os.File, n int) error {
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return err
	}

	size := info.Size()
	start := size - int64(n)
	if start < 0 {
		start = 0
	}

	_, err = f.Seek(start, 0)
	if err != nil {
		return err
	}

	buf := make([]byte, 4096)
	for {
		r, err := f.Read(buf)
		if r > 0 {
			_, _ = os.Stdout.Write(buf[:r])
		}
		if err != nil {
			break
		}
	}

	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: tail -c <num> file...")
		os.Exit(1)
	}

	if os.Args[1] != "-c" {
		fmt.Fprintln(os.Stderr, "only -c option is supported")
		os.Exit(1)
	}

	n, err := parsePositiveInt(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	files := os.Args[3:]
	multiple := len(files) > 1
	exitCode := 0
	printedAny := false // true if we've printed anything (error or file content)

	for _, path := range files {
		f, err := os.Open(path)
		if err != nil {
			// print error (no header), mark that we printed something
			fmt.Println(err)
			printedAny = true
			exitCode = 1
			continue
		}

		// opened successfully; print header if multiple files
		if multiple {
			if printedAny {
				fmt.Println()
			}
			fmt.Printf("==> %s <==\n", path)
		}

		// print tail content
		err = tailFromOpen(f, n)
		if err != nil {
			fmt.Println(err)
			exitCode = 1
		}

		printedAny = true
	}

	os.Exit(exitCode)
}
