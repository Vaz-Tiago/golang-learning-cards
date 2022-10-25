//lint:file-ignore U1000 Ignore all unused code, it's generated
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

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
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

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

func (d deck) toString() string {
	// that is possible because deck "extends" slice type
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile (filename string) deck {
	bs, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle () {
	// go rand always use the same seed to generate rand number
	// based on that, its necessary provide to rand a new seed (or source) to go use
	// on rand generation function
	rand.Seed(time.Now().UnixNano())
	
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}