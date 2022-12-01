package main

import "fmt"

func vyriesPostupnost() {
	postupnost := zasobnik{1, 2}
	for {
		if noveCislo(postupnost) <= 10000 {
			postupnost = append(postupnost, noveCislo(postupnost))
		} else {
			break
		}
	}

	fmt.Println(postupnost)
	fmt.Println("Pocet clenov slice je", len(postupnost))
}

func noveCislo(zasob zasobnik) int {
	vysledok := zasob[len(zasob)-1] + zasob[len(zasob)-2]
	return vysledok
}
