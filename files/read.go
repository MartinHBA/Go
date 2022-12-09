package main

import (
	"fmt"
	"os"
)

func readMyFile(myFileName string) string {
	file, err := os.Open(myFileName)
	if err != nil {
		fmt.Println(err)
		return "Error opening file!"
	}
	defer file.Close() // Close the file when the function returns

	// Read the file's contents into a byte slice
	data := make([]byte, 100)
	file.Read(data)
	return string(data)
}
