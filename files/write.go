package main

import (
	"fmt"
	"os"
)

func writeMyFile(filename string, veta string) {

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // Close the file when the function returns

	// Write some data to the file
	data := []byte(veta)
	n, err := file.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the number of bytes written and the contents of the file
	fmt.Println(n, string(data))

}
