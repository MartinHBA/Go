package main

type Zasobnik []int

func delejCisla() Zasobnik {
	Cisla := Zasobnik{}
	for i := 1; i < 11; i++ {
		Cisla = append(Cisla, i)
	}
	return Cisla
}
