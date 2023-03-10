package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	x := 5
	y := 10
	fmt.Printf("Before swapping: x = %d, y = %d\n", x, y)

	// Pass the addresses of x and y to the swap function
	swap(&x, &y)

	fmt.Printf("After swapping: x = %d, y = %d\n", x, y)
}
