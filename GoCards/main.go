package main

func main() {

	cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

	//	cards := newDeckFromFile("mycards")
	cards.shuffle()
	cards.print()

	// cards := newDeck()
	// cards.saveToFile("mycards")
	// fmt.Println(cards.toString())
	// fmt.Println([]byte(cards.toString()))
}
