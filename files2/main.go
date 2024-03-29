package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("strings.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

	fmt.Println("Strings from file:", lines)
}
