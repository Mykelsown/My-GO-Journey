package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}

	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}

	fileName := os.Args[1]

	content, err := os.ReadFile(fileName)
	if err != nil {
		return
	}

	fmt.Print(string(content))
}
