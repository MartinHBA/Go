package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().Unix())

	// Generate a random number between 1 and 100
	secretNumber := rand.Intn(100) + 1

	// Prompt the user to guess the number
	var guess int
	for {
		fmt.Print("Guess the number between 1 and 100: ")
		fmt.Scan(&guess)

		if guess > secretNumber {
			fmt.Println("Too high! Guess lower.")
		} else if guess < secretNumber {
			fmt.Println("Too low! Guess higher.")
		} else {
			fmt.Println("You guessed it! The secret number was", secretNumber)
			break
		}
	}
}
