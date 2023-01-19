package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Initialize the maze.
	maze := [][]rune{
		{'#', '#', '#', '#', '#'},
		{'#', '.', '.', '.', '#'},
		{'#', '#', '#', '.', '#'},
		{'#', '.', '.', '.', 'X'},
		{'#', '#', '#', '#', '#'},
	}
	// Initialize the starting position.
	row := 1
	col := 1
	// Keep track of the number of steps taken.
	steps := 0

	// Create a scanner to read input from the user.
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Print the current position in the maze.
		fmt.Println("You are currently at row", row, "column", col)

		// Check if the current position is the exit.
		if maze[row][col] == 'X' {
			fmt.Println("Congratulations, you found the exit!")
			fmt.Println("You took", steps, "steps.")
			break
		}

		// Print the possible moves.
		fmt.Println("Possible moves: w (up), s (down), a (left), d (right)")

		// Prompt the user for their next move.
		fmt.Print("Enter your next move: ")

		// Read the user's input.
		scanner.Scan()
		move := scanner.Text()

		// Check the move and update the position in the maze accordingly.
		switch move {
		case "w":
			// Check if the move is valid (not going out of bounds and not hitting a wall).
			if row > 0 && maze[row-1][col] != '#' {
				row--
				steps++
			} else {
				fmt.Println("Invalid move.")
			}
		case "s":
			if row < len(maze)-1 && maze[row+1][col] != '#' {
				row++
				steps++
			} else {
				fmt.Println("Invalid move.")
			}
		case "a":
			if col > 0 && maze[row][col-1] != '#' {
				col--
				steps++
			} else {
				fmt.Println("Invalid move.")
			}
		case "d":
			if col < len(maze[0])-1 && maze[row][col+1] != '#' {
				col++
				steps++
			} else {
				fmt.Println("Invalid move.")
			}
		}
	}
}
