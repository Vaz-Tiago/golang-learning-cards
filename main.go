package main

func main () {
	// cards := newDeck()
	// cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")
	cards.print()
	// hand, remainCards := deal(cards, 5)

	// hand.print()
	// fmt.Println("------------")
	// // remainCards.print()
	// fmt.Println(remainCards.toString())
}