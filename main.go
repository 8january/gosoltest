package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	_ "strings"
)

func main() {

	if len(os.Args) != 5 {
		fmt.Println("[USE] gosoltest <file.cpp> <input.in> <expected.sol> <output.sol>")
		os.Exit(1)
	}

	codeFile := os.Args[1]
	inputFile := os.Args[2]
	solutionFile := os.Args[3]
	outputFile := os.Args[4]

	fmt.Println("Compiling ", codeFile)
	cmd := exec.Command("g++", "-O2", "-Wall", codeFile, "-o", "a.out")
	err := cmd.Run()
	if err != nil {
		fmt.Println("[ERROR]: Compilation failed")
		fmt.Println("[ERROR]:", err)
		fmt.Println("[ERROR]: Check your C++ code.")
		os.Exit(1)
	}

	if _, err := os.Stat("a.out"); os.IsNotExist(err) {
		fmt.Println("[ERROR]:", err)
		fmt.Println("[ERROR]: Executable file not found.")
		os.Exit(1)
	}

	input, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("[ERROR]:", err)
	}

	cmd = exec.Command("./a.out")
	cmd.Stdin = strings.NewReader(string(input))
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("[ERROR]:", err)
	}

	err = os.WriteFile(outputFile, output, 0644)
	if err != nil {
		fmt.Println("[ERROR]:", err)
	}
	fmt.Println("output saved to ->", outputFile)

	expected, err := os.ReadFile(solutionFile)
	if err != nil {
		fmt.Println("[ERROR]:", err)
	}

	fmt.Printf("Comparing solutions...\n")
	if strings.TrimSpace(string(output)) == strings.TrimSpace(string(expected)) {
		fmt.Println("Test passed! the solution is correct.")
	} else {
		fmt.Println("Test failed! the solution is incorrect.")
		fmt.Println("Generated output:")
		fmt.Println(string(output))
		fmt.Println("Expected output:")
		fmt.Println(string(expected))
	}

}
