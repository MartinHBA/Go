package main

import "fmt"

func vyriesPostupnost() {
	postupnost := zasobnik{1, 2}
	for i := 1; i < 11; i++ {
		postupnost = append(postupnost, noveCislo(postupnost))
	}

	fmt.Println(postupnost)
	fmt.Println("Pocet clenov slice je", len(postupnost))
}

func noveCislo(zasob zasobnik) int {
	vysledok := zasob[len(zasob)-1] + zasob[len(zasob)-2]
	return vysledok
}
