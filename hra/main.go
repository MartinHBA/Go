package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Board [4][4]string
type move []int

func main() {
	newBoard := prepareBoard()
	printBoard(newBoard)
	var currentMove, counterMove move

	currentMove = getValidMove(newBoard)
	newBoard = changeBoard(newBoard, currentMove, 1)
	printBoard(newBoard)
	firstCounterMove := GetFirstCounterMove(newBoard)
	newBoard = changeBoard(newBoard, firstCounterMove, 2)
	printBoard(newBoard)
	for {
		// prerobit neskor podmienku ukoncenia
		if 1 == 0 {
			break
		} else {
			currentMove = getValidMove(newBoard)
			newBoard = changeBoard(newBoard, currentMove, 1)
			printBoard(newBoard)
			counterMove = getCounterMove(newBoard)
			if reflect.DeepEqual(counterMove, move{0, 0}) {
				fmt.Println("Game Over!")
				break
			}
			newBoard = changeBoard(newBoard, counterMove, 2)
			printBoard(newBoard)
		}
	}

}

func changeBoard(myBoard Board, myMove move, player int) Board {
	if player == 1 {
		myBoard[myMove[0]][myMove[1]] = "X"
		return myBoard
	} else {
		myBoard[myMove[0]][myMove[1]] = "O"
		return myBoard
	}

}

func getUserInput() move {
	fmt.Println("Enter row letter, comma and column number, something like 'B,2' and hit enter ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		indata := scanner.Text()
		if indata == "quit" {
			break
		}
		valuesSlice := strings.Split(indata, ",")
		valuesSlice[0] = strings.ToUpper(valuesSlice[0])
		intVar, err := strconv.Atoi(valuesSlice[1])
		if err != nil {
			fmt.Println("Error during conversion")
			return nil
		}
		myRow := ConvertLetterToInt(valuesSlice[0])
		return move{myRow, intVar}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)

	}
	return nil
}

func ConvertLetterToInt(letter string) int {
	if letter == "A" {
		return 1
	} else if letter == "B" {
		return 2
	} else if letter == "C" {
		return 3
	}
	return 0
}

func printBoard(myBoard Board) {
	fmt.Println(myBoard[0])
	fmt.Println(myBoard[1])
	fmt.Println(myBoard[2])
	fmt.Println(myBoard[3])
	fmt.Println("==========")
}

func prepareBoard() Board {
	var newBoard Board
	for i := 0; i < 4; i++ {
		newBoard[0][i] = strconv.Itoa(i)
		for z := 1; z < 4; z++ {

			newBoard[z][i] = " "
		}
	}
	newBoard[1][0] = "A"
	newBoard[2][0] = "B"
	newBoard[3][0] = "C"
	return newBoard
}
