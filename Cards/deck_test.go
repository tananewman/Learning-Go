package main

import "testing"

func TestGetNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Error("Expected deck length of 52 but got", len(d))
	}
}
