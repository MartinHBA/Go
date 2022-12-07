package main

import (
	"fmt"
	"strings"
)

const (
	boardSize = 3
)

var (
	board    = [boardSize][boardSize]string{}
	player   = "X"
	gameOver = false
)

func main() {
	for !gameOver {
		printBoard()
		fmt.Printf("Player %s, enter your move (row, col): ", player)
		var row, col int
		fmt.Scanf("%d,%d", &row, &col)
		if row >= boardSize || col >= boardSize {
			fmt.Println("Invalid move, please try again.")
			continue
		}
		if board[row][col] != "" {
			fmt.Println("That square is already occupied, please try again.")
			continue
		}
		board[row][col] = player
		if checkWin() {
			gameOver = true
			printBoard()
			fmt.Printf("Player %s wins!\n", player)
			break
		}
		if checkTie() {
			gameOver = true
			printBoard()
			fmt.Println("It's a tie!")
			break
		}
		player = "O"
	}
}

func printBoard() {
	for i := 0; i < boardSize; i++ {
		fmt.Println(strings.Join(board[i][:], " | "))
	}
}

func checkWin() bool {
	// check rows
	for i := 0; i < boardSize; i++ {
		if board[i][0] == player &&
			board[i][1] == player &&
			board[i][2] == player {
			return true
		}
	}
	// check cols
	for i := 0; i < boardSize; i++ {
		if board[0][i] == player &&
			board[1][i] == player &&
			board[2][i] == player {
			return true
		}
	}
	// check diagonals
	if board[0][0] == player &&
		board[1][1] == player &&
		board[2][2] == player {
		return true
	}
	if board[0][2] == player &&
		board[1][1] == player &&
		board[2][0] == player {
		return true
	}
	return false
}
func checkTie() bool {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == "" {
				// empty square found, game is not over
				return false
			}
		}
	}
	// no empty squares found, game is a tie
	return true
}
