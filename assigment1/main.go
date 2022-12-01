package main

import (
	"fmt"
)

func main() {

	radCisel := delejCisla()
	fmt.Println(radCisel)

	for _, value := range radCisel {
		if value%2 == 0 {
			fmt.Println(value, "is even")
		} else {
			fmt.Println(value, "is odd")
		}
	}
}
