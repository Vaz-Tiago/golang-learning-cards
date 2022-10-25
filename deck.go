package main

import "fmt"

// Create a new type of 'deck'
// witch is a slice of string
type deck []string

func (d deck) print(){
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four" } // "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}