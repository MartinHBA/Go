package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var words = []string{"holub", "zajac", "slon", "moriak", "leopard"}

func main() {
	rand.Seed(time.Now().UnixNano())
	word := words[rand.Intn(len(words))]
	fmt.Println("Vitaj v hre!")
	fmt.Println("Slovo je:", strings.Repeat("_", len(word)))
	var guessed string
	var incorrect int
	for {
		fmt.Print("Hadaj pismeno: ")
		var letter string
		fmt.Scan(&letter)
		if strings.Contains(guessed, letter) {
			fmt.Println("Toto pismeno si uz skusal, skus ine...")
			continue
		}
		guessed += letter
		if !strings.Contains(word, letter) {
			incorrect++
			fmt.Println("Nespravne. Mas este ", 6-incorrect, "pokusov.")
		} else {
			fmt.Println("Spravne!")
			display := ""
			for _, r := range word {
				if strings.Contains(guessed, string(r)) {
					display += string(r)
				} else {
					display += "_"
				}
			}
			fmt.Println("Slovo je:", display)
			if strings.Replace(display, "_", "", -1) == word {
				fmt.Println("VYHRA!")
				return
			}
		}
		if incorrect == 6 {
			fmt.Println("Prehra. Slovo bolo :", word)
			return
		}
	}
}
