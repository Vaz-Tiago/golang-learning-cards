package main

import (
	"os"
	"testing"
)

func TestNewDeck (t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("The first card of the deck must be Ace of Spades, but got %v", d[0])
	}

	if d[len(d) - 1] != "King of Clubs" {
		t.Errorf("The last card of deck must be King of Clubs, but got %v", d[len(d) -1])
	}
}

func TestDeal (t *testing.T) {
	cards := newDeck()

	hand, cards := deal(cards, 10)

	if len(hand) != 10 {
		t.Errorf("The deal function need give 10 cards, but got %v", len(hand))
	}

	if hand[0] != "Ace of Spades" {
		t.Errorf("The first card of hand must be 'Ace of Spades', but is '%v'", hand[0])
	}

	if hand[len(hand) - 1] != "Ten of Spades" {
		t.Errorf("The last card of hand must be 'Ten of Spade', but is '%v'", hand[len(hand) - 1])
	}

	if cards[0] != "Jack of Spade" {
		t.Errorf("The first card of deck after the deal must be 'Jack of Spade', but is '%v'", cards[0])
	}

	if len(cards) != 42 {
		t.Errorf("The deck must be 42 cards after deal, but got %v", len(cards))
	}
}

func TestShuffle (t *testing.T) {
	cards := newDeck()
	cards.shuffle()

	if cards[0] == "Ace of Spades" && cards[1] == "Two of Spades" {
		t.Errorf("The first two cards os deck cannot be Ace and Two of Spades")
	}
}

func TestSaveToFileAndNewDeckFromFile (t *testing.T) {
	os.Remove("_deckTesting")

	deck := newDeck()
	deck.saveToFile("_deckTesting")
	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expect 52 cards in loaded deck, got %v", len(loadedDeck))
	}

	os.Remove("_deckTesting")
}