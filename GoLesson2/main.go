package main

import "fmt"

type Board [3][3]string

type position struct {
	axisx int
	axisy int
	value string
}

func main() {

	myBoard := prepareBoard()
	fmt.Println(myBoard)
	// print po riadkoch
	printBoard(&myBoard)

	myMove := position{axisx: 1, axisy: 2, value: "O"}
	myBoard.setBoardValue(myMove)
	printBoard(&myBoard)

	myMove = position{axisx: 1, axisy: 1, value: "O"}
	myBoard.setBoardValue(myMove)
	printBoard(&myBoard)

	myMove = position{axisx: 1, axisy: 0, value: "O"}
	myBoard.setBoardValue(myMove)
	printBoard(&myBoard)
}

func prepareBoard() Board {
	var myBoard Board
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			myBoard[i][j] = "X"
		}
	}
	return myBoard
}

func printBoard(myBoard *Board) {
	fmt.Println("Here is current board:")
	fmt.Println(myBoard[0])
	fmt.Println(myBoard[1])
	fmt.Println(myBoard[2])
}

func (myBoard *Board) setBoardValue(myPosition position) {
	myBoard[myPosition.axisx][myPosition.axisy] = myPosition.value
}
