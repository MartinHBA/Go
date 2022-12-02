package main

import (
	"fmt"
	"os"
	"reflect"
)

func getDefensiveMove(myBoard Board) move {
	resultDiag := 0
	resultDiag2 := 0
	for i := 1; i < 4; i++ {
		resultRows := 0
		resultCols := 0
		z := 0
		if myBoard[i][i] == "X" {
			resultDiag++
			//fmt.Println(i, i, "diag -1- is ", resultDiag)
		}
		if myBoard[4-i][i] == "X" {
			resultDiag2++
			//fmt.Println(4-i, i, " diag -2- is ", resultDiag2)
		}
		for _, rowVal := range myBoard[i] {
			if rowVal == "X" {
				resultRows++
			}

			if myBoard[z][i] == "X" {
				resultCols++
			}

			z++
		}

		if resultRows == 2 {
			for q := 1; q < 4; q++ {
				if myBoard[i][q] == " " {
					return move{i, q}
				}
			}
		}

		//fmt.Println("Column ", i, " sum = ", resultCols)
		if resultCols == 2 {
			for q := 1; q < 4; q++ {
				if myBoard[q][i] == " " {
					return move{q, i}
				}
			}
		}

	}

	if resultDiag == 2 {
		//fmt.Println("Diagonal Detected Defensive move!")
		for i := 1; i < 4; i++ {
			if myBoard[i][i] == " " {
				return move{i, i}
			}
		}

	}

	if resultDiag2 == 2 {
		//fmt.Println("cross Diagonal Detected Defensive move!")
		for i := 1; i < 4; i++ {
			if myBoard[4-i][i] == " " {
				return move{4 - i, i}
			}
		}
	}

	return move{0, 0}
}

func getCounterMove(myBoard Board) move {
	var mymove move
	// check if winning move

	mymove = getWinningMove(myBoard)
	if reflect.DeepEqual(mymove, move{0, 0}) {
		// check defensive move
		mymove = getDefensiveMove(myBoard)
	}
	if reflect.DeepEqual(mymove, move{0, 0}) {
		mymove = randomMove(myBoard)
	}

	return mymove
}

func randomMove(myBoard Board) move {
	for i := 1; i < 4; i++ {
		for z := 1; z < 4; z++ {
			if myBoard[i][z] == " " {
				return move{i, z}
			}
		}
	}
	return move{0, 0}
}

func getWinningMove(myBoard Board) move {
	resultDiag := 0
	resultDiag2 := 0
	for i := 1; i < 4; i++ {
		resultRows := 0
		resultCols := 0
		z := 0

		if myBoard[i][i] == "O" {
			resultDiag++
			//fmt.Println(i, i, "diag -1- is ", resultDiag)
		}
		if myBoard[4-i][i] == "O" {
			resultDiag2++
			//fmt.Println(4-i, i, " diag -2- is ", resultDiag2)
		}

		for _, rowVal := range myBoard[i] {
			if rowVal == "O" {
				resultRows++
			}

			if myBoard[z][i] == "O" {
				resultCols++
			}
			z++
		}

		//fmt.Println("Row ", i, " sum = ", resultRows)
		if resultRows == 2 {
			for q := 1; q < 4; q++ {
				if myBoard[i][q] == " " {
					endGameBoard(myBoard, move{i, q})
				}
			}
		}

		//fmt.Println("Column ", i, " sum = ", resultCols)
		if resultCols == 2 {
			for q := 1; q < 4; q++ {
				if myBoard[q][i] == " " {
					endGameBoard(myBoard, move{q, i})
				}
			}
		}
		if resultDiag == 2 {
			//fmt.Println("Diagonal Detected Offensive move!")
			for i := 1; i < 4; i++ {
				if myBoard[i][i] == " " {
					endGameBoard(myBoard, move{i, i})
				}
			}

		}

		if resultDiag2 == 2 {
			//fmt.Println("cross Diagonal Detected Offensive move!")
			for i := 1; i < 4; i++ {
				if myBoard[4-i][i] == " " {
					endGameBoard(myBoard, move{4 - i, i})
				}
			}
		}
	}
	return move{0, 0}
}

func endGameBoard(myBoard Board, myMove move) {
	myBoard = changeBoard(myBoard, myMove, 2)
	printBoard(myBoard)
	fmt.Println("Game Over, you've lost!")
	os.Exit(0)
}

func GetFirstCounterMove(myBoard Board) move {

	if myBoard[2][2] == "X" {
		return move{3, 3}
	} else {
		return move{2, 2}
	}

}

func getValidMove(myBoard Board) move {
	potentialMove := getUserInput()
	IsValid := CheckValidMove(myBoard, potentialMove)
	if IsValid {
		return potentialMove
	} else {
		fmt.Println("Invalid move, let's try again...")
		printBoard(myBoard)
		return getValidMove(myBoard)
	}
}

func CheckValidMove(myBoard Board, myMove move) bool {
	if myBoard[myMove[0]][myMove[1]] == " " {
		return true
	} else {
		return false
	}

}
