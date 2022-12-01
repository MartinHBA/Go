package main

import "fmt"

func naucSlice() {
	//vytvorime premennu postupnost typu zasobnik a vlozime cisla 1,2,3
	postupnost := zasobnik{40, 60, 80}
	fmt.Println(postupnost)
	fmt.Println("0. clen slice je", postupnost[0])
	fmt.Println("1. clen slice je", postupnost[1])
	fmt.Println("2. clen slice je", postupnost[2])
	fmt.Println("pocet clenov v slice je", len(postupnost))
	fmt.Println("Takto sa odkazujem na posledneho clena v slice ->", postupnost[len(postupnost)-1])
}
