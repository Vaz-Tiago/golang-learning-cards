package main

import "fmt"

func main () {
	cards := newDeck()
	hand, remainCards := deal(cards, 5)

	hand.print()
	fmt.Println("------------")
	// remainCards.print()
	fmt.Println(remainCards.toString())
}