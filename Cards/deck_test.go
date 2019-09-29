package main

import (
	"os"
	"testing"
)

func TestGetNewDeck(t *testing.T) {
	d := newDeck()

	// there should be 52 cards in the deck
	if len(d) != 52 {
		t.Error("Expected deck length of 52 but got", len(d))
	}

	// first card should be Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Error("Expected first card to be Ace of Spades but got", d[0])
	}

	// last card should be King of Clubs
	if d[len(d)-1] != "King of Clubs" {
		t.Error("Expected last card to be King of Clubs")
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Error("Expected 52 cards in deck, but got", len(loadedDeck))
	}
}
