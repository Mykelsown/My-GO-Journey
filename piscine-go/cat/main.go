package main

import (
	"os"
)

func main() {
	// No arguments → read from stdin
	if len(os.Args) == 1 {
		buf := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(buf)
			if n > 0 {
				os.Stdout.Write(buf[:n])
			}
			if err != nil {
				return
			}
		}
	}

	// Arguments = filenames
	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			os.Stdout.Write([]byte("ERROR: " + err.Error() + "\n"))
			os.Exit(1)
		}

		buf := make([]byte, 1024)
		for {
			n, rErr := file.Read(buf)
			if n > 0 {
				os.Stdout.Write(buf[:n])
			}
			if rErr != nil {
				break
			}
		}

		file.Close()
	}
}
