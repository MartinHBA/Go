package main

import (
	"fmt"
)

func main() {
	// Open a file for reading
	readResult := readMyFile("file.txt")

	// Print the contents of the file
	fmt.Println(readResult)

	writeMyFile("file2.txt", "Ahoj ako sa mas?")

}
