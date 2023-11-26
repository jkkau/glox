package main

import (
	"bufio"
	"fmt"
	"os"

	"github.jkkau.glox/scanner"
)

func runFile(filename string) {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	run(string(fileData))
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		data, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading from stdin:", err)
			os.Exit(1)
		}
		
		run(string(data))
	}
}

func run(source string) {
	s := scanner.NewScanner(source)
	tokens := s.ScanTokens()
	for _, token := range tokens {
		fmt.Println(token)
	}
}

func main() {
	argsLen := len(os.Args)
	if argsLen > 2 {
		fmt.Println("Usage: go run main.go <filename>")
	}else if argsLen == 2 {
		filename := os.Args[1]
		runFile(filename)
	}else {
		runPrompt()
	}
}